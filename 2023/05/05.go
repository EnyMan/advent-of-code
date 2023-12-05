package fifth

import (
	helpers "advent/utils"
	"fmt"
	"math"
	"strings"
)

type SeedMapping struct {
	destination_range_starts []int
	source_range_starts      []int
	range_lengths            []int
	mapping_size             int
}

func (mc *SeedMapping) GetDestination(source int) int {
	for i := 0; i < mc.mapping_size; i++ {
		if mc.source_range_starts[i] <= source && mc.source_range_starts[i]+mc.range_lengths[i]-1 >= source {
			return mc.destination_range_starts[i] + (source - mc.source_range_starts[i])
		}
	}
	return source
}

func generate_mapping(line_splits [][]string) SeedMapping {
	destination_range_starts := []int{}
	source_range_starts := []int{}
	range_lengths := []int{}
	for _, line_split := range line_splits {
		destination_range_starts = append(destination_range_starts, helpers.ToInt(line_split[0]))
		source_range_starts = append(source_range_starts, helpers.ToInt(line_split[1]))
		range_lengths = append(range_lengths, helpers.ToInt(line_split[2]))
	}
	return SeedMapping{destination_range_starts, source_range_starts, range_lengths, len(destination_range_starts)}
}

func find_location(seed int,
	seed_to_soil *SeedMapping,
	soil_to_fertilizer *SeedMapping,
	fertilizer_to_water *SeedMapping,
	water_to_light *SeedMapping,
	light_to_temperature *SeedMapping,
	temperature_to_humidity *SeedMapping,
	humidity_to_location *SeedMapping,
) int {
	soil := (*seed_to_soil).GetDestination(seed)
	fertilizer := (*soil_to_fertilizer).GetDestination(soil)
	water := (*fertilizer_to_water).GetDestination(fertilizer)
	light := (*water_to_light).GetDestination(water)
	temperature := (*light_to_temperature).GetDestination(light)
	humidity := (*temperature_to_humidity).GetDestination(temperature)
	location := (*humidity_to_location).GetDestination(humidity)
	// fmt.Println(seed, "->", soil, "->", fertilizer, "->", water, "->", light, "->", temperature, "->", humidity, "->", location)
	return location
}

func parse_input(lines []string, part2 bool) int {
	valid_headers := []string{
		"seeds:", "seed-to-soil", "soil-to-fertilizer",
		"fertilizer-to-water", "water-to-light", "light-to-temperature",
		"temperature-to-humidity", "humidity-to-location",
	}
	seeds := []int{}
	header := ""
	lines_to_parse := map[string][][]string{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		// fmt.Println(line)
		line_split := strings.Fields(line)

		header_ := line_split[0]
		if header_ == "seeds:" {
			seeds = helpers.ArrayToInt(line_split[1:])
			continue
		}

		if helpers.FindInList(header_, valid_headers) {
			header = header_
			continue
		}

		lines_to_parse[header] = append(lines_to_parse[header], line_split)
	}

	seed_to_soil := generate_mapping(lines_to_parse["seed-to-soil"])
	soil_to_fertilizer := generate_mapping(lines_to_parse["soil-to-fertilizer"])
	fertilizer_to_water := generate_mapping(lines_to_parse["fertilizer-to-water"])
	water_to_light := generate_mapping(lines_to_parse["water-to-light"])
	light_to_temperature := generate_mapping(lines_to_parse["light-to-temperature"])
	temperature_to_humidity := generate_mapping(lines_to_parse["temperature-to-humidity"])
	humidity_to_location := generate_mapping(lines_to_parse["humidity-to-location"])

	min_location := math.MaxInt64

	for i := 0; i < len(seeds); i++ {
		fmt.Println(i, len(seeds))
		seed := 0
		if part2 {
			for j := 0; j < seeds[i+1]; j++ {
				seed = seeds[i] + j
				location := find_location(seed, &seed_to_soil, &soil_to_fertilizer, &fertilizer_to_water, &water_to_light, &light_to_temperature, &temperature_to_humidity, &humidity_to_location)
				if location < min_location {
					min_location = location
				}
			}
			i++
		} else {
			seed = seeds[i]
			location := find_location(seed, &seed_to_soil, &soil_to_fertilizer, &fertilizer_to_water, &water_to_light, &light_to_temperature, &temperature_to_humidity, &humidity_to_location)
			if location < min_location {
				min_location = location
			}
		}
	}
	return min_location
}

func Fifth(filename string, part2 bool) {
	lines := helpers.ReadFile(filename)
	min_location := parse_input(lines, part2)
	fmt.Println(min_location)
}
