using DataStructures;
f = open("2022/05_input.txt", "r")
input = readlines(f)
stacks = Dict{Int, Deque{String}}()
stacks9001 = Dict{Int, Deque{String}}()
stack_count = ceil(length(input[1])/4)
header = true

for line in input
    global header
    if line == ""
        header = false
        continue
    end
    if header
        for stack in 1:stack_count
            if !haskey(stacks, Int(stack))
                stacks[Int(stack)] = Deque{String}()
                stacks9001[Int(stack)] = Deque{String}()
            end
            offset = Int(4 * (stack - 1))
            crate = line[begin + offset:begin + offset + 2]
            if !contains(crate, ']')
                continue
            end 
            push!(stacks[Int(stack)], crate)
            push!(stacks9001[Int(stack)], crate)
        end
    else
        line = replace(line, "move" => "")
        line = replace(line, "from" => "")
        line = replace(line, "to" => "")
        count, from, to = map(x -> parse(Int, x), split(line))
        # @show count, from, to
        crane = Deque{String}()
        for n in 1:count
            # @show first(stacks[from])
            crate = popfirst!(stacks[from])
            pushfirst!(stacks[to], crate)

            crate = popfirst!(stacks9001[from])
            pushfirst!(crane, crate)
        end
        for crate in crane
            pushfirst!(stacks9001[to], crate)
        end
    end
end

for stack in 1:stack_count
    stack = Int(stack)
    # @show stack, stacks[stack]
    crate = replace(first(stacks[stack]), "]" => "")
    crate = replace(crate, "[" => "")
    print(crate)
end
println("")
for stack in 1:stack_count
    stack = Int(stack)
    # @show stack, stacks[stack]
    crate = replace(first(stacks9001[stack]), "]" => "")
    crate = replace(crate, "[" => "")
    print(crate)
end
println("")
