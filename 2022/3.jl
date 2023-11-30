f = open("03_input.txt", "r")
int(x) = floor(Int, x)
backpacks = readlines(f)


group_score = []
current_group = Vector{String}()
score = []
for backpack in backpacks
    global current_group
    @show length(current_group)
    # @show current_group

    if length(current_group) < 3
        push!(current_group, String(backpack))
    end
    if length(current_group) == 3
        @show current_group
        found = false
        for item in current_group[1]
            if occursin(item, current_group[2]) && occursin(item, current_group[3]) && !found
                @show item
                if islowercase(item)
                    append!(group_score, Int(item) - 96)
                else
                    append!(group_score, Int(item) - 65 + 27)
                end
                found = true
            end
        end
        empty!(current_group)
    end

    # @show length(backpack)
    found = false
    middle = int(length(backpack)/2)
    first_half = backpack[begin:end-middle]
    second_half = backpack[end-middle+1:end]
    # unique(intersect(first_half, second_half))
    for item in second_half
        if occursin(item, first_half) && !found
            if islowercase(item)
                # @show backpack, first_half, second_half, item
                append!(score, Int(item) - 96)
            else
                # @show backpack, first_half, second_half, item
                append!(score, Int(item) - 65 + 27)
            end
            found = true
        end
    end
end
@show sum(score)
@show sum(group_score)