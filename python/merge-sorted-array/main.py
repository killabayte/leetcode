def merge(nums1: List[int], m: int, nums2: List[int], n: int) -> None:
    index1 = m-1
    index2 = n-1
    indexMerge = m+n-1

    while index2>=0 and index1>=0:
        if nums1[index1]>nums2[index2]:
            nums1[indexMerge] = nums1[index1]
            index1 -= 1
        else:
            nums1[indexMerge] = nums2[index2]
            index2 -= 1
        
        indexMerge -= 1
    
    while index2>=0:
        nums1[indexMerge] = nums2[index2]
        index2 -= 1
        indexMerge -= 1

merge([1,2,3,0,0,0], 3, [2,5,6], 3)