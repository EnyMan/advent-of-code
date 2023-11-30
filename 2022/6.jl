using DataStructures;
f = open("2022/06_input.txt", "r")
input = readline(f)
header = Vector{Char}()
processed = String("")

for char in input
    global processed
    before_push = length(Set(header))
    push!(header, char)
    processed *= char
    after_push = length(Set(header))
    if after_push > before_push && after_push == 14
        @show header, processed, length(processed)
        break
    elseif after_push == before_push
        index = findfirst(x -> x==char, header)
        for i in Base.OneTo(index)
            popfirst!(header)
        end
    end
end