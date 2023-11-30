f = open("04_input.txt", "r")
pairs = readlines(f)

full_overlap = 0
partial_overlap = 0
for pair in pairs
    elf1, elf2 = split(pair,',')
    elf1_start, elf1_end = map(x -> parse(Int, x), split(elf1, '-'))
    elf2_start, elf2_end = map(x -> parse(Int, x), split(elf2, '-'))
    elf1_range = elf1_start:elf1_end
    elf2_range = elf2_start:elf2_end

    intersection = intersect(Set(elf1_range), Set(elf2_range))
    if intersection == Set(elf1_range) || intersection == Set(elf2_range)
        global full_overlap += 1
    end
    if length(intersection) > 0
        global partial_overlap += 1
    end
end
@show full_overlap
@show partial_overlap
