func lemonadeChange(bills []int) bool {
    fives, tens := 0, 0

    for i := 0; i < len(bills); i++ {
        curBill := bills[i]

        switch curBill {
            case 5:
                fives++
            case 10:
                if fives >= 1 {
                    fives--
                    tens++
                } else {
                    return false
                }
            case 20:
                if tens >= 1 && fives >= 1 {
                    tens--
                    fives--
                } else if fives >= 3 {
                    fives -= 3
                } else {
                    return false
                }
        }
    }

    return true
}