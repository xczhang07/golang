func isPalindrome(x int) bool {
    var copy_x int = x
    var r_x int = 0
    for x > 0 {
        tmp := x % 10
        r_x = r_x * 10 + tmp
        x = x/10
    }
    if r_x == copy_x {
        return true
    }
    return false
}
