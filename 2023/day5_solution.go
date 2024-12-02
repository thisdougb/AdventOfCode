package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Data struct {
	Seeds                 []int64 `json:"seeds"`
	SeedToSoil            []Map   `json:"seed_to_soil"`
	SoilToFertilizer      []Map   `json:"soil_to_fertilizer"`
	FertilizerToWater     []Map   `json:"fertilizer_to_water"`
	WaterToLight          []Map   `json:"water_to_light"`
	LightToTemperature    []Map   `json:"light_to_temperataure"`
	TemperatureToHumidity []Map   `json:"temperature_to_humidity"`
	HumidityToLocation    []Map   `json:"humidity_to_location"`
}
type Map struct {
	DestinationRangeStart int64 `json:"destination_range_start"`
	SourceRangeStart      int64 `json:"source_range_start"`
	RangeLength           int64 `json:"range_length"`
}

type Seed struct {
	seed        int64
	soil        int64
	fertilizer  int64
	water       int64
	light       int64
	temperature int64
	humidity    int64
	location    int64
}

func mapsTo(m []Map, source int64) int64 {

	for _, n := range m {

		if source < n.SourceRangeStart || source >= n.SourceRangeStart+n.RangeLength {
			continue
		}

		offset := source - n.SourceRangeStart
		//fmt.Printf("seed %d : source range %d-%d offset %d \n", source, n.SourceRangeStart, n.SourceRangeStart+n.RangeLength, offset)

		return n.DestinationRangeStart + offset
	}

	return source
}

var (
	data Data
)

func day5() {
	day5part2("day5_input.txt")
}

/*
(py3.9) dougb $ go run ./...
lowest location 47909639
*/
func day5part2(filepath string) {

	input := readFileContents(filepath)
	loadData(input, true)

	lowest := int64(-1)
	for _, s := range data.Seeds {

		soil := mapsTo(data.SeedToSoil, s)
		fertilizer := mapsTo(data.SoilToFertilizer, soil)
		water := mapsTo(data.FertilizerToWater, fertilizer)
		light := mapsTo(data.WaterToLight, water)
		temperature := mapsTo(data.LightToTemperature, light)
		humidity := mapsTo(data.TemperatureToHumidity, temperature)
		location := mapsTo(data.HumidityToLocation, humidity)

		if lowest == -1 || location < lowest {
			lowest = location
		}
	}

	fmt.Printf("lowest location %d\n", lowest)
}

/*
(py3.9) dougb $ go run ./...
seed 3037945983, soil 3617161702, fertilizer 4143012697, water 3935232580, light 4037900502, temperature 3383604305, humidity 477256891, location 2555693727
seed 743948277, soil 923133255, fertilizer 664480320, water 2187923114, light 1402085698, temperature 2151358085, humidity 1103305405, location 2716845820
seed 2623786093, soil 3951552909, fertilizer 3632313287, water 3744317805, light 2990053102, temperature 2796793130, humidity 3390855216, location 3421775940
seed 391282324, soil 1210472639, fertilizer 951819704, water 72361177, light 3468277158, temperature 3275017186, humidity 368669772, location 1326688836
seed 195281306, soil 326406338, fertilizer 511649069, water 1172526594, light 2413333744, temperature 2413333744, humidity 1277377578, location 3482819330
seed 62641412, soil 1033278867, fertilizer 774625932, water 2298068726, light 576360916, temperature 435008621, humidity 1600918510, location 4034180178
seed 769611781, soil 948796759, fertilizer 690143824, water 2213586618, light 491878808, temperature 350526513, humidity 1516436402, location 3949698070
seed 377903357, soil 1197093672, fertilizer 938440737, water 58982210, light 3454898191, temperature 3261638219, humidity 355290805, location 1313309869
seed 2392990228, soil 3720757044, fertilizer 4246608039, water 3657086698, light 3254536665, temperature 3061276693, humidity 3655338779, location 1526086962
seed 144218002, soil 1114855457, fertilizer 856202522, water 2379645316, light 657937506, temperature 516585211, humidity 1682495100, location 4115756768
seed 1179463071, soil 1775935261, fertilizer 1517282326, water 192638979, light 3588554960, temperature 2483908223, humidity 1347952057, location 3553393809
seed 45174621, soil 1015812076, fertilizer 757159141, water 2280601935, light 558894125, temperature 417541830, humidity 1583451719, location 4016713387
seed 2129467491, soil 4103933688, fertilizer 3784694066, water 3800463992, light 3046199289, temperature 2852939317, humidity 3447001403, location 3080164134
seed 226193957, soil 357318989, fertilizer 2157776697, water 696297733, light 4290588219, temperature 3636292022, humidity 4222063225, location 226172555
seed 1994898626, soil 1994898626, fertilizer 1736245691, water 2456495242, light 734787432, temperature 593435137, humidity 1227853320, location 2841393735
seed 92402726, soil 1063040181, fertilizer 804387246, water 2327830040, light 606122230, temperature 464769935, humidity 1630679824, location 4063941492
seed 1555863421, soil 259766145, fertilizer 445008876, water 2143078857, light 1357241441, temperature 2106513828, humidity 1058461148, location 2672001563
seed 340215202, soil 1159405517, fertilizer 900752582, water 21294055, light 21294055, temperature 1709262720, humidity 2815470407, location 1277267466
seed 426882817, soil 615000806, fertilizer 2415458514, water 1330444910, light 2555180678, temperature 3743214442, humidity 259988683, location 622612797
seed 207194644, soil 338319676, fertilizer 523562407, water 1184439932, light 2425247082, temperature 2425247082, humidity 1289290916, location 3494732668
lowest location 226172555
*/
func day5part1(filepath string) {

	seeds := []Seed{}

	input := readFileContents(filepath)
	loadData(input, false)

	for _, s := range data.Seeds {

		soil := mapsTo(data.SeedToSoil, s)
		fertilizer := mapsTo(data.SoilToFertilizer, soil)
		water := mapsTo(data.FertilizerToWater, fertilizer)
		light := mapsTo(data.WaterToLight, water)
		temperature := mapsTo(data.LightToTemperature, light)
		humidity := mapsTo(data.TemperatureToHumidity, temperature)
		location := mapsTo(data.HumidityToLocation, humidity)

		seed := Seed{s, soil, fertilizer, water, light, temperature, humidity, location}
		seeds = append(seeds, seed)
	}

	lowest := seeds[0].location
	for _, seed := range seeds {

		if seed.location < lowest {
			lowest = seed.location
		}

		fmt.Printf("seed %d, soil %d, fertilizer %d, water %d, light %d, temperature %d, humidity %d, location %d \n",
			seed.seed,
			seed.soil,
			seed.fertilizer,
			seed.water,
			seed.light,
			seed.temperature,
			seed.humidity,
			seed.location)
	}
	fmt.Printf("lowest location %d\n", lowest)
}

func loadData(input []string, p2 bool) {

	for i := 0; i < len(input); i++ {
		if strings.HasPrefix(input[i], "seeds:") {
			if p2 {
				parts := strings.Split(input[i], " ")
				for p := 1; p < len(parts); p += 2 {
					seedStart, _ := strconv.ParseInt(parts[p], 10, 64)
					seedLength, _ := strconv.ParseInt(parts[p+1], 10, 64)

					for i := seedStart; i < seedStart+seedLength; i++ {
						data.Seeds = append(data.Seeds, i)
					}
				}

			} else {
				parts := strings.Split(input[i], " ")
				for p := 1; p < len(parts); p++ {
					seed, _ := strconv.ParseInt(parts[p], 10, 64)
					data.Seeds = append(data.Seeds, seed)
				}
			}
		}
		if input[i] == "seed-to-soil map:" {
			data.SeedToSoil = extractMapValues(&input, i)
			i += len(data.SeedToSoil)
		}
		if input[i] == "fertilizer-to-water map:" {
			data.FertilizerToWater = extractMapValues(&input, i)
			i += len(data.FertilizerToWater)
		}
		if input[i] == "soil-to-fertilizer map:" {
			data.SoilToFertilizer = extractMapValues(&input, i)
			i += len(data.SoilToFertilizer)
		}
		if input[i] == "water-to-light map:" {
			data.WaterToLight = extractMapValues(&input, i)
			i += len(data.WaterToLight)
		}
		if input[i] == "light-to-temperature map:" {
			data.LightToTemperature = extractMapValues(&input, i)
			i += len(data.LightToTemperature)
		}
		if input[i] == "temperature-to-humidity map:" {
			data.TemperatureToHumidity = extractMapValues(&input, i)
			i += len(data.TemperatureToHumidity)
		}
		if input[i] == "humidity-to-location map:" {
			data.HumidityToLocation = extractMapValues(&input, i)
			i += len(data.HumidityToLocation)
		}
		if i == len(input)-1 {
			break
		}
	}
}

func extractMapValues(input *[]string, pos int) []Map {

	values := []Map{}
	pos++
	for {
		if (*input)[pos] == "" || pos == len(*input)-1 {
			break
		}

		mapvalues := stringToint64Array((*input)[pos])
		m := Map{mapvalues[0], mapvalues[1], mapvalues[2]}

		values = append(values, m)
		pos++
	}
	return values
}

func stringToint64Array(input string) []int64 {
	a := []int64{}

	parts := strings.Split(input, " ")
	for _, p := range parts {
		i, _ := strconv.ParseInt(p, 10, 64)
		a = append(a, i)
	}
	return a
}
