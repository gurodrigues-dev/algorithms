array = list(range(1,10000))

def binarySearch(array, target):

    smaller = 0
    bigger = len(array) - 1
    
    while smaller <= bigger:
    
        mid = (smaller + bigger) // 2
        guess = array[mid]
    
        if guess == target:
            return mid
    
        if guess > target:
            bigger = mid - 1
    
        if guess < target:
            smaller = mid + 1

    return None

print(binarySearch(array, 9532))