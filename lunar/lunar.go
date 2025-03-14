package lunar

import (
	"fmt"
	"strings"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/animal"
	"github.com/Lofanmi/chinese-calendar-golang/utils"
)

// Lunar 农历
type Lunar struct {
	t                *time.Time
	year, month, day int64
	monthIsLeap      bool
}

var numberAlias = [...]string{
	"零", "一", "二", "三", "四",
	"五", "六", "七", "八", "九",
}

var dateAlias = [...]string{
	"初", "十", "廿", "卅",
}

var lunarMonthAlias = [...]string{
	"正", "二", "三", "四", "五", "六",
	"七", "八", "九", "十", "冬", "腊",
}

var lunars = [...]int64{
	0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2, // 1900-1909
	0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977, // 1910-1919
	0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970, // 1920-1929
	0x06566, 0x0d4a0, 0x0ea50, 0x16a95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950, // 1930-1939
	0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557, // 1940-1949
	0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5b0, 0x14573, 0x052b0, 0x0a9a8, 0x0e950, 0x06aa0, // 1950-1959
	0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0, // 1960-1969
	0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b6a0, 0x195a6, // 1970-1979
	0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570, // 1980-1989
	0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x05ac0, 0x0ab60, 0x096d5, 0x092e0, // 1990-1999
	0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5, // 2000-2009
	0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930, // 2010-2019
	0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530, // 2020-2029
	0x05aa0, 0x076a3, 0x096d0, 0x04afb, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45, // 2030-2039
	0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0, // 2040-2049
	0x14b63, 0x09370, 0x049f8, 0x04970, 0x064b0, 0x168a6, 0x0ea50, 0x06aa0, 0x1a6c4, 0x0aae0, // 2050-2059
	0x092e0, 0x0d2e3, 0x0c960, 0x0d557, 0x0d4a0, 0x0da50, 0x05d55, 0x056a0, 0x0a6d0, 0x055d4, // 2060-2069
	0x052d0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x0b6a6, 0x0ad50, 0x055a0, 0x0aba4, 0x0a5b0, 0x052b0, // 2070-2079
	0x0b273, 0x06930, 0x07337, 0x06aa0, 0x0ad50, 0x14b55, 0x04b60, 0x0a570, 0x054e4, 0x0d160, // 2080-2089
	0x0e968, 0x0d520, 0x0daa0, 0x16aa6, 0x056d0, 0x04ae0, 0x0a9d4, 0x0a2d0, 0x0d150, 0x0f252, // 2090-2099
	0x0d520, 0x0db27, 0x0b5a0, 0x055d0, 0x04db5, 0x049b0, 0x0a4b0, 0x0d4b4, 0x0aa50, 0x0b559, // 2100-2109
	0x06d20, 0x0ad60, 0x05766, 0x09370, 0x04970, 0x06974, 0x054b0, 0x06a50, 0x07a53, 0x06aa0, // 2110-2119
	0x1aaa7, 0x0aad0, 0x052e0, 0x0cae5, 0x0a960, 0x0d4a0, 0x1e4a4, 0x0d950, 0x05abb, 0x056a0, // 2120-2129
	0x0a6d0, 0x151d6, 0x052d0, 0x0a8d0, 0x1d155, 0x0b2a0, 0x0b550, 0x06d52, 0x055a0, 0x1a5a7, // 2130-2139
	0x0a5b0, 0x052b0, 0x0a975, 0x068b0, 0x07290, 0x0baa4, 0x06b50, 0x02dbb, 0x04b60, 0x0a570, // 2140-2149
	0x052e6, 0x0d160, 0x0e8b0, 0x06d25, 0x0da90, 0x05b50, 0x036d3, 0x02ae0, 0x0a3d7, 0x0a2d0, // 2150-2159
	0x0d150, 0x0d556, 0x0b520, 0x0d690, 0x155a4, 0x055b0, 0x02afa, 0x045b0, 0x0a2b0, 0x0aab6, // 2160-2169
	0x0a950, 0x0b4a0, 0x1b2a5, 0x0ad50, 0x055b0, 0x02b73, 0x04570, 0x06377, 0x052b0, 0x06950, // 2170-2179
	0x06d56, 0x05aa0, 0x0ab50, 0x056d4, 0x04ae0, 0x0a570, 0x06562, 0x0d2a0, 0x0eaa6, 0x0d550, // 2180-2189
	0x05aa0, 0x0aea5, 0x0a6d0, 0x04ae0, 0x0aab3, 0x0a4d0, 0x0d2b7, 0x0b290, 0x0b550, 0x15556, // 2190-2199
	0x02da0, 0x095d0, 0x145b4, 0x049b0, 0x0a4f9, 0x064b0, 0x06a90, 0x0b696, 0x06b50, 0x02b60, // 2200-2209
	0x09b64, 0x09370, 0x04970, 0x06963, 0x0e4a0, 0x0eaa7, 0x0da90, 0x05b50, 0x02ed5, 0x02ae0, // 2210-2219
	0x092e0, 0x1c2d4, 0x0c950, 0x0d4d9, 0x0b4a0, 0x0b690, 0x057a7, 0x055b0, 0x025d0, 0x095b5, // 2220-2229
	0x092b0, 0x0a950, 0x1c953, 0x0b4a0, 0x0b5a8, 0x0ad50, 0x055b0, 0x12375, 0x02570, 0x052b0, // 2230-2239
	0x1a2b4, 0x06950, 0x06cbb, 0x05aa0, 0x0ab50, 0x14ad6, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, // 2240-2249
	0x0e950, 0x07553, 0x05aa0, 0x0aba7, 0x095d0, 0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x0da55, // 2250-2259
	0x0b540, 0x0d6a0, 0x0ada1, 0x095b0, 0x04b77, 0x049b0, 0x0a4b0, 0x0b4b5, 0x06a50, 0x0ad40, // 2260-2269
	0x1ab53, 0x02b60, 0x19568, 0x09370, 0x04970, 0x06566, 0x0e4a0, 0x0ea50, 0x16a94, 0x05ad0, // 2270-2279
	0x02b60, 0x0aae2, 0x092e0, 0x0cad6, 0x0c950, 0x0d4a0, 0x0dca5, 0x0b650, 0x056a0, 0x0b5b3, // 2280-2289
	0x025d0, 0x093b7, 0x092b0, 0x0a950, 0x0b556, 0x074a0, 0x0b550, 0x05d54, 0x04da0, 0x0a5b0, // 2290-2299
	0x06572, 0x052b0, 0x0aaa6, 0x0e950, 0x06aa0, 0x1aaa5, 0x0ab50, 0x04b60, 0x0aae3, 0x0a570, // 2300-2309
	0x052d7, 0x0d260, 0x0d950, 0x16956, 0x056a0, 0x09ad0, 0x145d4, 0x04ad0, 0x0a4fa, 0x0a4d0, // 2310-2319
	0x0d250, 0x1d457, 0x0b540, 0x0b6a0, 0x195a5, 0x095b0, 0x049b0, 0x0a973, 0x0a4b0, 0x0b2b8, // 2320-2329
	0x06a50, 0x06d40, 0x0b746, 0x0ab60, 0x09570, 0x142f4, 0x04970, 0x064b0, 0x074a3, 0x0ea50, // 2330-2339
	0x16c57, 0x05ac0, 0x0ab60, 0x096d5, 0x092e0, 0x0c960, 0x0d954, 0x0d4a0, 0x0daa8, 0x0b550, // 2340-2349
	0x056a0, 0x1a9b6, 0x025d0, 0x092d0, 0x0cab5, 0x0a950, 0x0b4a0, 0x0f4a1, 0x0b550, 0x15557, // 2350-2359
	0x04ba0, 0x0a5b0, 0x05575, 0x052b0, 0x0a930, 0x07954, 0x06aa0, 0x0ada8, 0x0ab50, 0x04b60, // 2360-2369
	0x0a6e6, 0x0a570, 0x05260, 0x0ea65, 0x0d920, 0x0daa0, 0x156a2, 0x096d0, 0x04bd7, 0x04ad0, // 2370-2379
	0x0a4d0, 0x0d4b5, 0x0d250, 0x0d520, 0x1d544, 0x0b5a0, 0x056ea, 0x095b0, 0x049b0, 0x0a576, // 2380-2389
	0x0a4b0, 0x0b250, 0x0ba54, 0x06d20, 0x0ada0, 0x06b62, 0x09370, 0x04af6, 0x04970, 0x064b0, // 2390-2399
	0x06ca5, 0x0ea50, 0x06b20, 0x0bac3, 0x0ab60, 0x093d8, 0x092e0, 0x0c960, 0x0d556, 0x0d4a0, // 2400-2409
	0x0da50, 0x05d55, 0x056a0, 0x0aad0, 0x065d2, 0x052d0, 0x1a8b7, 0x0a950, 0x0b4a0, 0x1b2a5, // 2410-2419
	0x0ad50, 0x055a0, 0x0aba3, 0x0a5b0, 0x15278, 0x05270, 0x06930, 0x07536, 0x06aa0, 0x0ad50, // 2420-2429
	0x14b54, 0x04b60, 0x0a570, 0x144e3, 0x0d260, 0x1e867, 0x0d520, 0x0da90, 0x06ea5, 0x056d0, // 2430-2439
	0x04ae0, 0x0a9d4, 0x0a4d0, 0x0d2b8, 0x0d250, 0x0d520, 0x0db27, 0x0b5a0, 0x056d0, 0x04db5, // 2440-2449
	0x049b0, 0x0a4b0, 0x1c4b3, 0x0aa50, 0x0b558, 0x06d20, 0x0ad60, 0x15365, 0x05370, 0x04970, // 2450-2459
	0x06974, 0x064b0, 0x06aa8, 0x0ea50, 0x06aa0, 0x1aaa6, 0x0aad0, 0x052e0, 0x0cae5, 0x0c960, // 2460-2469
	0x0d4a0, 0x0f4a3, 0x0d950, 0x05b57, 0x056a0, 0x0a6d0, 0x055d5, 0x052d0, 0x0a950, 0x0d954, // 2470-2479
	0x0b4a0, 0x0b56a, 0x0ad50, 0x055a0, 0x0a7a6, 0x0a5b0, 0x052b0, 0x0a975, 0x06930, 0x07290, // 2480-2489
	0x1aa93, 0x06d50, 0x12d57, 0x04b60, 0x0a570, 0x052e5, 0x0d160, 0x0e8b0, 0x16524, 0x0da90, // 2490-2499
	0x06b6a, 0x056d0, 0x02ae0, 0x0a5d6, 0x0a2d0, 0x0d150, 0x1d155, 0x0b520, 0x0da90, 0x075a2, // 2500-2509
	0x055b0, 0x02bb7, 0x045b0, 0x0a2b0, 0x0b2b5, 0x0a950, 0x0b520, 0x0bd24, 0x0ad50, 0x055b0, // 2510-2519
	0x05371, 0x04570, 0x16176, 0x052b0, 0x06950, 0x16955, 0x05aa0, 0x0ab50, 0x14ad3, 0x04ae0, // 2520-2529
	0x1a4e7, 0x0a560, 0x0d4a0, 0x0eaa6, 0x0d950, 0x05aa0, 0x1a6a4, 0x0a6d0, 0x04ae0, 0x0cab1, // 2530-2539
	0x0a8d0, 0x0d4b7, 0x0b290, 0x0b550, 0x15555, 0x035a0, 0x095d0, 0x055b3, 0x049b0, 0x0a977, // 2540-2549
	0x068b0, 0x06a90, 0x0b696, 0x06b50, 0x02da0, 0x09b64, 0x09570, 0x051e8, 0x0d160, 0x0e4a0, // 2550-2559
	0x0eaa7, 0x0da90, 0x05b50, 0x02ed5, 0x02ae0, 0x092e0, 0x0d2d4, 0x0c950, 0x0d557, 0x0b4a0, // 2560-2569
	0x0b690, 0x15996, 0x055b0, 0x029d0, 0x095b4, 0x0a2b0, 0x1a939, 0x0a950, 0x0b4a0, 0x0b6a6, // 2570-2579
	0x0ad50, 0x055a0, 0x0ab74, 0x02570, 0x052b0, 0x0b2b3, 0x06950, 0x06d57, 0x05aa0, 0x0ab50, // 2580-2589
	0x056d5, 0x04ae0, 0x0a570, 0x05554, 0x0d260, 0x0e96a, 0x0d550, 0x05aa0, 0x1aaa7, 0x096d0, // 2590-2599
	0x04ae0, 0x1a1b5, 0x0a4d0, 0x0d250, 0x1d253, 0x0b540, 0x1d658, 0x02da0, 0x095b0, 0x14976, // 2600-2609
	0x049b0, 0x0a4b0, 0x0b4b4, 0x06a50, 0x0b55b, 0x06b50, 0x02b60, 0x09766, 0x09370, 0x04970, // 2610-2619
	0x16165, 0x0e4a0, 0x0ea50, 0x07a93, 0x05ac0, 0x0abd8, 0x02ae0, 0x092e0, 0x0cad6, 0x0c950, // 2620-2629
	0x0d4a0, 0x0dca5, 0x0b650, 0x056a0, 0x0d5b1, 0x025d0, 0x093b7, 0x092b0, 0x0a950, 0x1d155, // 2630-2639
	0x074a0, 0x0b550, 0x14d53, 0x055a0, 0x1a568, 0x0a570, 0x052b0, 0x0aaa6, 0x0e950, 0x06ca0, // 2640-2649
	0x1aaa4, 0x0ab50, 0x04b60, 0x18ae2, 0x0a570, 0x052d7, 0x0d260, 0x0e920, 0x0ed55, 0x05aa0, // 2650-2659
	0x09ad0, 0x056d3, 0x04ad0, 0x0a5b7, 0x0a4d0, 0x0d250, 0x0da56, 0x0b540, 0x0d6a0, 0x09da4, // 2660-2669
	0x095b0, 0x04ab0, 0x0a973, 0x0a4b0, 0x0b2b7, 0x06a50, 0x06d40, 0x1b345, 0x0ab60, 0x095b0, // 2670-2679
	0x05373, 0x04970, 0x06567, 0x0d4a0, 0x0ea50, 0x06e56, 0x05ac0, 0x0ab60, 0x096d4, 0x092e0, // 2680-2689
	0x0c960, 0x0e953, 0x0d4a0, 0x0daa7, 0x0b550, 0x056a0, 0x0ada5, 0x0a5d0, 0x092d0, 0x0d2b3, // 2690-2699
	0x0a950, 0x1b458, 0x074a0, 0x0b550, 0x15556, 0x04da0, 0x0a5b0, 0x05574, 0x052b0, 0x0a930, // 2700-2709
	0x16933, 0x06aa0, 0x1aca7, 0x0ab50, 0x04b60, 0x1a2e5, 0x0a560, 0x0d260, 0x1e264, 0x0d920, // 2710-2719
	0x0dac9, 0x0d6a0, 0x09ad0, 0x149d6, 0x04ad0, 0x0a4d0, 0x0d4b5, 0x0d250, 0x0d53b, 0x0b540, // 2720-2729
	0x0b6a0, 0x057a7, 0x095b0, 0x049b0, 0x1a175, 0x0a4b0, 0x0b250, 0x0ba54, 0x06d20, 0x0adc9, // 2730-2739
	0x0ab60, 0x09570, 0x04af6, 0x04970, 0x064b0, 0x06ca5, 0x0ea50, 0x06d20, 0x19aa2, 0x0ab50, // 2740-2749
	0x152d7, 0x092e0, 0x0c960, 0x0d556, 0x0d4a0, 0x0da50, 0x15554, 0x056a0, 0x1aaa8, 0x0a5d0, // 2750-2759
	0x052d0, 0x0aab6, 0x0a950, 0x0b4a0, 0x1b4a5, 0x0b550, 0x055a0, 0x0aba3, 0x0a5b0, 0x05377, // 2760-2769
	0x05270, 0x06930, 0x07536, 0x06aa0, 0x0ad50, 0x05b53, 0x04b60, 0x0a5e8, 0x0a4e0, 0x0d260, // 2770-2779
	0x0ea66, 0x0d520, 0x0da90, 0x06ea5, 0x056d0, 0x04ae0, 0x0aad3, 0x0a4d0, 0x0d2b7, 0x0d250, // 2780-2789
	0x0d520, 0x1d926, 0x0b6a0, 0x056d0, 0x055b3, 0x049b0, 0x1a478, 0x0a4b0, 0x0aa50, 0x0b656, // 2790-2799
	0x06d20, 0x0ad60, 0x05b64, 0x05370, 0x04970, 0x06973, 0x064b0, 0x06aa7, 0x0ea50, 0x06b20, // 2800-2809
	0x0aea6, 0x0ab50, 0x05360, 0x1c2e4, 0x0c960, 0x0d4d9, 0x0d4a0, 0x0da50, 0x05b57, 0x056a0, // 2810-2819
	0x0a6d0, 0x055d5, 0x052d0, 0x0a950, 0x1c953, 0x0b490, 0x0b5a8, 0x0ad50, 0x055a0, 0x1a3a5, // 2820-2829
	0x0a5b0, 0x052b0, 0x1a174, 0x06930, 0x072b9, 0x06a90, 0x06d50, 0x02f56, 0x04b60, 0x0a570, // 2830-2839
	0x054e5, 0x0d160, 0x0e920, 0x0f523, 0x0da90, 0x06ba8, 0x056d0, 0x02ae0, 0x0a5d6, 0x0a4d0, // 2840-2849
	0x0d150, 0x0d955, 0x0d520, 0x0daa9, 0x0b590, 0x056b0, 0x02bb7, 0x049b0, 0x0a2b0, 0x0b2b5, // 2850-2859
	0x0aa50, 0x0b520, 0x1ad23, 0x0ad50, 0x15567, 0x05370, 0x04970, 0x06576, 0x054b0, 0x06a50, // 2860-2869
	0x07954, 0x06aa0, 0x0ab6a, 0x0aad0, 0x05360, 0x0a6e6, 0x0a960, 0x0d4a0, 0x0eca5, 0x0d950, // 2870-2879
	0x05aa0, 0x0b6a3, 0x0a6d0, 0x04bd7, 0x04ab0, 0x0a8d0, 0x0d4b6, 0x0b290, 0x0b540, 0x0dd54, // 2880-2889
	0x055a0, 0x095ea, 0x095b0, 0x052b0, 0x0a976, 0x068b0, 0x07290, 0x1b295, 0x06d50, 0x02da0, // 2890-2899
	0x18b63, 0x09570, 0x150e7, 0x0d160, 0x0e8a0, 0x1e8a6, 0x0da90, 0x05b50, 0x126d4, 0x02ae0, // 2900-2909
	0x092fb, 0x0a2d0, 0x0d150, 0x0d557, 0x0b4a0, 0x0da90, 0x05d95, 0x055b0, 0x02ad0, 0x185b3, // 2910-2919
	0x0a2b0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x1b4a6, 0x0ad50, 0x055a0, 0x0ab64, 0x0a570, 0x052f9, // 2920-2929
	0x052b0, 0x06950, 0x06d57, 0x05aa0, 0x0ab50, 0x152d5, 0x04ae0, 0x0a570, 0x05554, 0x0d260, // 2930-2939
	0x0e9a8, 0x0d950, 0x05aa0, 0x1aaa6, 0x096d0, 0x04ad0, 0x0aab4, 0x0a4d0, 0x0d2b8, 0x0b250, // 2940-2949
	0x0b540, 0x0d757, 0x02da0, 0x095b0, 0x04db5, 0x049b0, 0x0a4b0, 0x0b4b4, 0x06a50, 0x0b598, // 2950-2959
	0x06d50, 0x02d60, 0x09766, 0x09370, 0x04970, 0x06964, 0x0e4a0, 0x0ea6a, 0x0da50, 0x05b40, // 2960-2969
	0x1aad7, 0x02ae0, 0x092e0, 0x0cad5, 0x0c950, 0x0d4a0, 0x1d4a3, 0x0b650, 0x15658, 0x055b0, // 2970-2979
	0x029d0, 0x191b6, 0x092b0, 0x0a950, 0x0d954, 0x0b4a0, 0x0b56a, 0x0ad50, 0x055a0, 0x0a766, // 2980-2989
	0x0a570, 0x052b0, 0x0aaa5, 0x0e950, 0x06ca0, 0x0baa3, 0x0ab50, 0x04bd8, 0x04ae0, 0x0a570, // 2990-2999
	0x150d6, // 3000
}

// NewLunar 创建农历对象
func NewLunar(t *time.Time) *Lunar {
	year, month, day, isLeap := FromSolarTimestamp(t.Unix())
	return &Lunar{
		t:           t,
		year:        year,
		month:       month,
		day:         day,
		monthIsLeap: isLeap,
	}
}

// FromSolarTimestamp 通过时间戳创建
func FromSolarTimestamp(ts int64) (lunarYear, lunarMonth, lunarDay int64, lunarMonthIsLeap bool) {
	var (
		i, offset, leap         int64
		daysOfYear, daysOfMonth int64
		isLeap                  bool
	)
	// 与 1900-01-31 相差多少天
	t := time.Unix(ts, 0)
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	t2 := time.Date(1900, 1, 31, 0, 0, 0, 0, time.UTC)
	offset = (t1.Unix() - t2.Unix()) / 86400

	for i = 1900; i <= 3000 && offset > 0; i++ {
		daysOfYear = daysOfLunarYear(i)
		offset -= daysOfYear
	}
	if offset < 0 {
		offset += daysOfYear
		i--
	}

	// 农历年
	lunarYear = i
	// 闰哪个月
	leap = leapMonth(i)

	isLeap = false

	// 用当年的天数 offset, 逐个减去每月(农历)的天数, 求出当天是本月的第几天
	for i = 1; i < 13 && offset > 0; i++ {
		// 闰月
		if leap > 0 && i == (leap+1) && !isLeap {
			i--
			isLeap = true
			// 计算农历月天数
			daysOfMonth = leapDays(lunarYear)
		} else {
			// 计算农历普通月天数
			daysOfMonth = lunarDays(lunarYear, i)
		}
		// 解除闰月
		if isLeap && i == (leap+1) {
			isLeap = false
		}
		offset -= daysOfMonth
	}
	// offset 为 0 时, 并且刚才计算的月份是闰月, 要校正
	if offset == 0 && leap > 0 && i == leap+1 {
		if isLeap {
			isLeap = false
		} else {
			isLeap = true
			i--
		}
	}
	if offset < 0 {
		offset += daysOfMonth
		i--
	}
	// 农历月
	lunarMonth = i
	// 农历日
	lunarDay = offset + 1
	// 农历是否为闰月
	lunarMonthIsLeap = isLeap

	return
}

// ToSolarTimestamp 转换为时间戳
func ToSolarTimestamp(year, month, day, hour, minute, second int64, isLeapMonth bool) int64 {
	var (
		i, offset int64
	)
	// 参数合法性效验
	if year < 1900 || year > 3000 {
		return 0
	}
	// 参数区间 1900.1.31~3000.12.1
	m := leapMonth(year)
	// 传参要求计算该闰月公历 但该年得出的闰月与传参的月份并不同
	if isLeapMonth && (m != month) {
		isLeapMonth = false
	}
	// 超出了最大极限值
	if year == 3000 && month == 12 && day > 1 || year == 1900 && month == 1 && day < 31 {
		return 0
	}
	days := lunarDays(year, month)
	maxDays := days
	// if month is leap, _day use leapDays method
	if isLeapMonth {
		maxDays = leapDays(year)
	}
	// 参数合法性效验
	if day > maxDays {
		return 0
	}
	// 计算农历的时间差
	offset = 0
	for i = 1900; i < year; i++ {
		offset += daysOfLunarYear(i)
	}
	isAdd := false
	for i = 1; i < month; i++ {
		leap := leapMonth(year)
		if !isAdd {
			// 处理闰月
			if leap <= i && leap > 0 {
				offset += leapDays(year)
				isAdd = true
			}
		}
		offset += lunarDays(year, i)
	}
	// 转换闰月农历 需补充该年闰月的前一个月的时差
	if isLeapMonth {
		offset += days
	}
	// 1900 年农历正月初一的公历时间为 1900年1月31日0时0分0秒 (该时间也是本农历的最开始起始点)
	var startTimestamp int64 = -2203804800
	t1 := time.Unix((offset+day-31)*86400+startTimestamp, 0)
	t2 := time.Date(t1.Year(), t1.Month(), t1.Day(), int(hour), int(minute), int(second), 0, time.Local)
	return t2.Unix()
}

// LeapMonth 获取闰月(0表示不闰, 5表示闰五月)
func (lunar *Lunar) LeapMonth() int64 {
	return leapMonth(lunar.year)
}

// IsLeap 是否闰年
func (lunar *Lunar) IsLeap() bool {
	return lunar.LeapMonth() != 0
}

// IsLeapMonth 是否闰月
func (lunar *Lunar) IsLeapMonth() bool {
	return lunar.monthIsLeap
}

// Animal 返回年份生肖
func (lunar *Lunar) Animal() *animal.Animal {
	return animal.NewAnimal(utils.OrderMod(lunar.year-3, 12))
}

// YearAlias 汉字表示年(二零一八)
func (lunar *Lunar) YearAlias() string {
	s := fmt.Sprintf("%d", lunar.year)
	for i, replace := range numberAlias {
		s = strings.Replace(s, fmt.Sprintf("%d", i), replace, -1)
	}
	return s
}

// MonthAlias 汉字表示月(八月, 闰六月)
func (lunar *Lunar) MonthAlias() string {
	pre := ""
	if lunar.monthIsLeap {
		pre = "闰"
	}
	return pre + lunarMonthAlias[lunar.month-1] + "月"
}

// DayAlias 汉字表示日(初一, 初十...)
func (lunar *Lunar) DayAlias() (alias string) {
	switch lunar.day {
	case 10:
		alias = "初十"
	case 20:
		alias = "二十"
	case 30:
		alias = "三十"
	default:
		alias = dateAlias[(int)(lunar.day/10)] + numberAlias[lunar.day%10]
	}
	return
}

// GetYear 年
func (lunar *Lunar) GetYear() int64 {
	return lunar.year
}

// GetMonth 月
func (lunar *Lunar) GetMonth() int64 {
	return lunar.month
}

// GetDay 日
func (lunar *Lunar) GetDay() int64 {
	return lunar.day
}

// Equals 返回两个对象是否相同
func (lunar *Lunar) Equals(b *Lunar) bool {
	return lunar.GetYear() == b.GetYear() &&
		lunar.GetMonth() == b.GetMonth() &&
		lunar.GetDay() == b.GetDay() &&
		lunar.IsLeapMonth() == b.IsLeapMonth()
}

func daysOfLunarYear(year int64) int64 {
	var (
		i, sum int64
	)
	sum = 29 * 12
	for i = 0x8000; i > 0x8; i >>= 1 {
		if (lunars[year-1900] & i) != 0 {
			sum++
		}
	}
	return sum + leapDays(year)
}

func leapMonth(year int64) int64 {
	return lunars[year-1900] & 0xf
}

func leapDays(year int64) (days int64) {
	if leapMonth(year) == 0 {
		days = 0
	} else if (lunars[year-1900] & 0x10000) != 0 {
		days = 30
	} else {
		days = 29
	}
	return
}

func lunarDays(year, month int64) (days int64) {
	if month > 12 || month < 1 {
		days = -1
	} else if (lunars[year-1900] & (0x10000 >> uint64(month))) != 0 {
		days = 30
	} else {
		days = 29
	}
	return
}
