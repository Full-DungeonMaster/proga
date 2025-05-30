from collections import deque

def read_maze(filename):
    with open(filename, 'r') as file:
        maze = [list(line.strip()) for line in file]
    return maze

def find_start_end(maze):
    start = end = None
    for i in range(len(maze)):
        for j in range(len(maze[i])):
            if maze[i][j] == 'S':
                start = (i, j)
            elif maze[i][j] == 'E':
                end = (i, j)
    return start, end

def bfs(maze, start, end):
    queue = deque()
    queue.append(start)
    visited = {start: None}

    while queue:
        current = queue.popleft()
        if current == end:
            break

        for direction in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
            neighbor = (current[0] + direction[0], current[1] + direction[1])
            if 0 <= neighbor[0] < len(maze) and 0 <= neighbor[1] < len(maze[0]):
                if maze[neighbor[0]][neighbor[1]] != '#' and neighbor not in visited:
                    queue.append(neighbor)
                    visited[neighbor] = current

    path = []
    current = end
    while current is not None:
        path.append(current)
        current = visited[current]
    path.reverse()
    return path

def write_solution(maze, path, filename):
    for (i, j) in path:
        if maze[i][j] not in ('S', 'E'):
            maze[i][j] = '0'
    with open(filename, 'w') as file:
        for row in maze:
            file.write(''.join(row) + '\n')

def solve_maze(input_file, output_file):
    maze = read_maze(input_file)
    start, end = find_start_end(maze)
    if not start or not end:
        raise ValueError("Лабиринт должен содержать начальную ('S') и конечную ('E') точки.")
    path = bfs(maze, start, end)
    write_solution(maze, path, output_file)