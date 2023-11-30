f = open("2022/02_input.txt", "r")
strategy = readlines(f)

points = Dict("X" => 1, "Y" => 2, "Z" => 3)

strength = Dict("A" => "Z", "B" => "X", "C" => "Y")
weakness = Dict("A" => "Y", "B" => "Z", "C" => "X")
equal = Dict("A"=>"X", "B" => "Y", "C" => "Z")

function result!(oponent, me)
    if strength[oponent] == me
        return "lose"
    elseif equal[oponent] == me
        return "draw"
    else
        return "win"
    end
end


function win(oponent)
    return weakness[oponent]
end

function draw(oponent)
    return equal[oponent]
end

function lose(oponent)
    return strength[oponent]
end

score = []
for round in strategy
    oponent, outcome = split(round, " ")
    if outcome == "Z"
        me = win(oponent)
        append!(score, points[me] + 6)
    elseif outcome == "Y"
        me = draw(oponent)
        append!(score, points[me] + 3)
    else
        me = lose(oponent)
        append!(score, points[me])
    end
end 

@show sum(score)