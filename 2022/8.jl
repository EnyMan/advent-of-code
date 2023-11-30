function parse_input()
    f = open("2022/08_input.txt", "r")
    input = readlines(f)
    trees = parse.(Int, split(input[1], ""))

    for i=2:lastindex(input)
        trees = hcat(trees, parse.(Int, split(input[i], "")))
    end
    return trees
end

function is_visible(tree, section)
    if length(section) == 0
        return true
    end
    for sub_tree in section
        if tree <= sub_tree
            return false
        end
    end
    return true
end

function can_see(tree, section)
    distance = 0
    if length(section) == 0
        return distance
    end
    for sub_tree in section
        if tree > sub_tree
            distance += 1
        elseif tree <= sub_tree 
            distance += 1
            return distance
        end
    end
    return distance
end

trees = parse_input()

visible = 0
max_distance = 0
for x in axes(trees, 1)
    for y in axes(trees, 2)
        global visible
        global max_distance
        tree = trees[x, y]
        # println("[$x, $y] = $tree")

        up = trees[x, begin:y]
        pop!(up)
        down = trees[x, y:end]
        popfirst!(down)
        left = trees[begin:x, y]
        pop!(left)
        right = trees[x:end, y]
        popfirst!(right)

        left_visible = is_visible(tree, left)
        right_visible =  is_visible(tree, right)
        up_visible =  is_visible(tree, up)
        down_visible =  is_visible(tree, down)

        left_distance = can_see(tree, reverse(left))
        right_distance =  can_see(tree, right)
        up_distance =  can_see(tree, reverse(up))
        down_distance =  can_see(tree, down)

        distance = up_distance * left_distance * down_distance * right_distance
        if distance > max_distance
            max_distance = distance
        end

        if left_visible | right_visible | up_visible | down_visible
            visible += 1
        end
    end
end

@show visible
@show max_distance