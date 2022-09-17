// https://www.hackerrank.com/challenges/bigger-is-greater

package main

import "fmt"

func biggerIsGreater(w string) string {
    bs := []byte(w)
    runningIdx := len(bs) - 1
    for runningIdx > 0 {
        char := bs[runningIdx]
        prevChar := bs[runningIdx - 1]
        if prevChar < char {
            insertFirstIntoSortedRest(bs[runningIdx-1:])
            reverseBytes(bs[runningIdx:])
            return string(bs)
        }
        runningIdx--
    }

    return "no answer"
}

// Scenarios:
// bs before: []byte{5, 10, 9, 8, 7, 6, 4, 3, 2, 1}
// bs after: []byte{6, 10, 9, 8, 7, 5, 4, 3, 2, 1}

// bs before: []byte{5, 10, 9, 8, 7, 6, 6, 3, 2, 1}
// bs after: []byte{6, 10, 9, 8, 7, 6, 5, 3, 2, 1}

// bs before: []byte{5, 10, 9, 8, 7, 5, 5, 3, 2, 1}
// bs after: []byte{7, 10, 9, 8, 5, 5, 5, 3, 2, 1}
func insertFirstIntoSortedRest(bs []byte) {
    if len(bs) < 2 {
        return
    }

    first := bs[0]  // 1

    for i := 1; i < len(bs); i++ {
        val := bs[i]
        if val <= first {
            bs[0], bs[i-1] = bs[i-1], bs[0]
            return
        }
    }

    bs[0], bs[len(bs)-1] = bs[len(bs)-1], bs[0]
}

func reverseBytes(bs []byte) {
    leftIdx := 0
    rightIdx := len(bs) - 1
    for leftIdx < rightIdx {
        bs[leftIdx], bs[rightIdx] = bs[rightIdx], bs[leftIdx]

        leftIdx++
        rightIdx--
    }
}

func main() {
    input := []string {
        "zedawdvyyfumwpupuinbdbfndyehircmylbaowuptgmw",
        "zyyxwwtrrnmlggfeb",
        "ocsmerkgidvddsazqxjbqlrrxcotrnfvtnlutlfcafdlwiismslaytqdbvlmcpapfbmzxmftrkkqvkpflxpezzapllerxyzlcf",
        "biehzcmjckznhwrfgglverxsezxuqpj",
        "rebjvsszebhehuojrkkhszxltyqfdvayusylgmgkdivzlpmmtvbsavxvydldmsym",
        "unpzhmbgrrs",
        "jprfovzkdlmdcesdcpdchcwoedjchcovklhrhlzfeeptoewcqpxg",
        "ywsfmynmiylcjgrfrrmtyeeykffzkuphpajndwxjteyjba",
        "dkuashjzsdq",
        "gwakhcpkolybihkmxyecrdhsvycjrljajlmlqgpcnmvvkjlkvdowzdfikh",
        "nebsajjbbuifimjpdcqfygeitief",
        "qetpicxagjkydehfnvfxrtigljlheulcsfidjjozbsnomygqbcmpffwswptbgkzrbgqwnczrcfynjmhebfbgseuhckbtuynvbo",
        "xuqfobndhsnsztifmqpnencxkllnpmbfqihtgdgxjhsvitlgtodhcydfb",
        "xqdwkjpkmrvkfnztozzlqtuxzxyxwowf",
        "yewluyxiwiprnajrtkeilolkmqidazhiar",
        "zzyyxxxxxwwwwwvvvvutttttsssssrrrrqqqppponnnnmmmmllkkjjjjiiggffffffeedddddbbbbbba",
        "hlvpzsfwnzsazeyaxaczkkrstiilkldupsqmzjnnsyoao",
        "zxvuutttrrrpoookiihhgggfdca",
    }

    for _, v := range input {
        fmt.Println(biggerIsGreater(v))
    }
}
