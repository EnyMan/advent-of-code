f = open("2022/01_input.txt", "r")
input = readlines(f)

elfs = Vector{Int64}()

elf = 1
for calories in input
    if calories == ""
        global elf += 1
        continue
    end
    if length(elfs) < elf        
        append!(elfs, 0)
    end
    elfs[elf] += parse(Int64, calories)
end
sorted_elfs = sort(elfs)
most_calories = last(sorted_elfs)
@show most_calories

@show sum(sorted_elfs[end-2:end])