import JSON
f = open("2022/07_input.txt", "r")
input = readlines(f)

# fix when there are multiple folder with the same name in different subdirectories
tree = Dict{String, Dict}()
current_dir = "/"
current_path = "/"

function calculate_dir_size(dirname)
    total_size = 0
    for subfolder in tree[dirname]["folders"]
        total_size += calculate_dir_size(subfolder)
    end
    for file in tree[dirname]["files"]
        total_size += file.second
    end
    return total_size
end

for line in input
    global tree
    global current_dir
    global current_path
    if contains(line, "cd") && contains(line, '$')
        marker, command, destination = split(line)
        if destination == ".."
            current_path = tree[current_path]["parent"] 
        else
            dir = Dict("name" => destination, "parent" => current_path, "files" => [], "folders" => [], "total_size" => 0)
            current_dir = destination
            if current_path == "/" && destination == "/"
                current_path = "/"
            elseif current_path == "/"
                current_path *= destination
            else
                current_path *= '/' * destination
            end
            tree[current_path] = dir
        end
    elseif contains(line, "dir")
        # subdir
        marker, subdir = split(line)
        if current_path == "/"
            push!(tree[current_path]["folders"], '/' * subdir)
        else
            push!(tree[current_path]["folders"], current_path * '/' * subdir)
        end
        
    else
        # files
        if contains(line, '$')
            continue
        end
        size, name = split(line)
        push!(tree[current_path]["files"], Pair(name, parse(Int, size)))
    end
end

sum_of_folders = 0
for dir in keys(tree)
    global sum_of_folders
    tree[dir]["total_size"] = calculate_dir_size(dir)
    if tree[dir]["total_size"] <= 100000
        # println(tree[dir])
        sum_of_folders += tree[dir]["total_size"]
    end
end

println(JSON.json(tree))
@show sum_of_folders

space_needed = 30000000 - 70000000 + tree["/"]["total_size"]
folder_to_delete = Dict("name" => "/", "parent" => "/", "files" => [], "folders" => [], "total_size" => 70000000)
@show space_needed

for dir in keys(tree)
    global folder_to_delete
    if tree[dir]["total_size"] >= space_needed && tree[dir]["total_size"] <= folder_to_delete["total_size"]
        folder_to_delete = tree[dir]
    end
end
@show folder_to_delete