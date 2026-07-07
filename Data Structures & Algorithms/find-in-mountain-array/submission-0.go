func findInMountainArray(target int, mountainArr *MountainArray) int {
    n := mountainArr.length()

    // Find peak
    lo, hi := 0, n-1
    for lo < hi {
        mid := lo + (hi-lo)/2
        if mountainArr.get(mid) < mountainArr.get(mid+1) {
            lo = mid + 1
        } else {
            hi = mid
        }
    }

    peak := lo

    // Search increasing part
    lo, hi = 0, peak
    for lo <= hi {
        mid := lo + (hi-lo)/2
        v := mountainArr.get(mid)

        if v == target {
            return mid
        }

        if v < target {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }

    // Search decreasing part
    lo, hi = peak+1, n-1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        v := mountainArr.get(mid)

        if v == target {
            return mid
        }

        if v < target {
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }

    return -1
}