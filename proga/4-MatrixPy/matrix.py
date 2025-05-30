def opred(matrix):
    n = len(matrix)
    if n == 1:
        return matrix[0][0]
    if n == 2:
        return matrix[0][0] * matrix[1][1] - matrix[0][1] * matrix[1][0]
    
    det = 0
    for col in range(n):
        minor = [row[:col] + row[col+1:] for row in matrix[1:]]
        det += ((-1) ** col) * matrix[0][col] * opred(minor)
    return det

def trace(matrix):
    return sum(matrix[i][i] for i in range(len(matrix)))

def transpose(matrix):
    return [[matrix[j][i] for j in range(len(matrix))] for i in range(len(matrix[0]))]



# Пример использования
mat = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
]

print("Определитель:", opred(mat))
print("След:", trace(mat))
print("Транспонированная матрица:")
for row in transpose(mat):
    print(row)