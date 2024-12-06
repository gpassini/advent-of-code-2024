package main

import (
	"fmt"
	"strings"
)

func main() {
	res := 0
	m := parseInput(input)
	// fmt.Println(m)
	for y, line := range m {
		for x := range line {
			res += tryAllDirs(m, x, y)
		}
	}
	fmt.Println(res)
}

func parseInput(in string) lettersMap {
	m := make(lettersMap, 0)
	lines := strings.Split(in, "\n")
	for _, l := range lines {
		runes := make([]rune, 0)
		for _, r := range l {
			runes = append(runes, r)
		}
		m = append(m, runes)
	}
	return m
}

func tryAllDirs(m lettersMap, x, y int) int {
	if m[y][x] != 'X' {
		return 0
	}

	res := 0
	var expectedRunes = []rune{'M', 'A', 'S'}
	for _, dir := range allDirections {
		if tryDir(m, x, y, dir, expectedRunes) {
			// fmt.Println("Found (", x, ",", y, ")", dir)
			res++
		}
	}
	return res
}

func tryDir(m lettersMap, x, y int, dir direction, searching []rune) bool {
	if len(searching) == 0 {
		// it's a match!
		return true
	}

	if found, newX, newY, ok := m.move(x, y, dir); !ok {
		// out-of-bounds
		return false
	} else if found != searching[0] {
		// not a match
		return false
	} else {
		// keep moving in that direction
		return tryDir(m, newX, newY, dir, searching[1:])
	}
}

type lettersMap [][]rune

// move returns the rune found after moving from the given origin in the given direction.
// The second return value, a boolean, will be false if moving out of bounds.
// The convention is:
// - x: est-west
// - y: north-south
// - the origin (0,0) is on the top left corner (or northwest)
// - lettersMap is assumed to be a rectangle (all lines have equal lengths, and all columns have equal lengths).
func (m lettersMap) move(x, y int, dir direction) (rune, int, int, bool) {
	switch dir {
	case north:
		y = y - 1
	case south:
		y = y + 1
	case east:
		x = x + 1
	case west:
		x = x - 1
	case northEst:
		y = y - 1
		x = x + 1
	case northWest:
		y = y - 1
		x = x - 1
	case southEst:
		y = y + 1
		x = x + 1
	case southWest:
		y = y + 1
		x = x - 1
	default:
		panic(fmt.Sprintf("unknown direction: %+v", dir))
	}
	if x < 0 || y < 0 || x >= len(m[0]) || y >= len(m) {
		return -1, x, y, false
	} else {
		return m[y][x], x, y, true
	}
}

func (m lettersMap) String() string {
	sb := strings.Builder{}
	for _, l := range m {
		for _, r := range l {
			sb.WriteRune(r)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

type direction uint

func (d direction) String() string {
	switch d {
	case north:
		return "N"
	case south:
		return "S"
	case east:
		return "E"
	case west:
		return "W"
	case northEst:
		return "NE"
	case northWest:
		return "NW"
	case southEst:
		return "SE"
	case southWest:
		return "SW"
	default:
		panic(fmt.Sprintf("unknown direction: %d", d))
	}
}

const (
	north direction = iota
	south
	east
	west
	northEst
	northWest
	southEst
	southWest
)

var (
	allDirections = []direction{
		north,
		south,
		east,
		west,
		northEst,
		northWest,
		southEst,
		southWest,
	}
)

// expected: 18
const sample = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

const input = `XXAMXSSMSSXMAXMSXXMAMMMXXXXSAMXMSMSMXSAMXMSMSXSXSSMSMXXSXSXSXSAXXMASASMSMXAASXXMAXXXXXXMXXAMSAMXSMSMMSXXMASMMSMMMMMMXMAXSXMMSXXAMXXXMMMMSAMX
SSMMXMAAAXMMSSMMMMMAAMSMAAAMXMAXMAAAMSAMXXAASAMAXAASXMASXMASAMAMSSMXASAAXSAMXMAXMMSSMMMSMSMAAAAAXMXAAXXXXAXMASAAAAAAAXSXMAXAAAXMSAMXMAMXMXAM
XAAMMSMMMSMAAAXMAXMASMAXSAMXASASMSMSMSAMXSMXMMMAMMMMAMXMASAMAMAMMAMMMMMMMAXXASMMAAAAAAAAAAXMSMMMSMAMMSMMMXSMASXSSSSSSSMASMMMSSXXAMAMSASXMSSS
SSMMAAMAAMMMMSMSXSAAXXMMXMAXMXMMAMAMMMMXXASXAMSMSMXSXMMSAMASXMAMSAMXSAXMAMMMXAAXMMSXXMXASXXXXMAMAXXAAAAAAAXMXSAXAMAMXAMAMAAXMAMXAMAXMASXAAAS
XMASXSSSSXAMSMAMASMXSMXMASMSSSMMAMAMMMMXMAMXMMAAAAMMAXAMMXAMXXSAMXMAMMSXMXASMSSXSXXMMSAXAMXMMSMSMMSMSSSMMSSMMMXMMMSSSMMMSMSXMAXSXMXSMAMMMMSM
ASXMXXAXXXMSAMSMASXMMXAXXMAAAAAMSSSMAAMSMAMASXMSMSASAMXSASAMXMMMSMMASMSASMMSAXXMMAXAASXMAAAMAMXAAMSXXAAXXXAMSXMXXAAXAASXAAXXMAMMSAMXMAXASMMM
XMMSSMMMSXXSASXMASAMAASMMMMMSMMMMAMSXMXAXAMMSAAXAXMAMSAMMAXXAXAMAMSASMSAMMSMXMAAMSMMMMXAAXMSSSSMSMSMMMMSXSAMAASAMXXXXMMSMSMMMMXMAXAXSMSMXAAA
SXASAAXAMMMXXMXMAMMMMMXAMXXXMMXSMAMAMSMMMSMXSMMMMMSMSMXSMMMSXXXXAMMAXAMAMSAAXSXSAXASXMSSXSXAAXAAMAXAAAXSASAMXXMASMMXAAAXXAAAXMASXMMMSXAASMMM
AMSSMMMXXAAMMXXMASXMXXSAMSSMSSXASMSSMSASAXMMXAAXXAMXAMXXAXAMXMSSMSMSMSMSXMMMMXAAASXMAAAAASMMSMMMMAMSSSSMAMAMAXMXMAASMMMSSSSMMSAXAAXAXXAMXAAA
MSMSXXAASMSMMAMMAMAXAAMAMAAAAMMAMXMAASAMASXMASMXMASMMSASMMSXAXMASAAAAXXAXSXMSMXMXMASMMMMMMAXXXXAMAMMAMXMMMMMASXMMMMMAAAAMAMXMMASMXMSSMSXSSMM
XAAMXMAMAAAASASMMSAMXSSSMMMMASAMXSSMMMAMAMMSMAMXMAXMXMAAXAMXSXSSSXSMSMSMAAMXAAMXMSASMXAXXXXMAMMXSXSXSAMXMASMAMAAMMMMXMMSMMMXXMMMMAAAAXAAMMXS
SMSMAASMSSSMSASAMAMXAMXXAXXSAMXSAAXAASAMXMAXXAMSMSMAXMSMMAXAMMSAMAXXAAAXSMMMMXMAAMAXXXMXSMMAMASAXAMXMASXSMSMMSXMMASXXMXMAMMSMSMXSXMASMMSMAAM
XMAXAXAMXAAAMAMMMXXMASMSXMMAXXAMXSSSMSMSSMASMSSXAXAMXMAXXAMSXMXAMXMSMSMMMXMASASASMSMSMAASAMMSASXSSMMSAMMAMXAMSAMSXMAMSAMXSAAAAAMAAXXAXXAMMSS
MSMSXSSMMSMMMMMXAXSSXMASMMSSMMXXMAAAAXAAXMAXAXAMXMSSXSASMXSAAMSSMSAAMAMXMASASASXXAAASMSMSAMMMASAMMAXMASMSMMSMMAMMAXAXSXMMAXMSMSASMMSMMSASAAM
XAASAMXAAXXXMAXMMSXAAXAXAAAAASASAMSMMMMMSXMMSMMXAAXXXMASMXMXSMAAASMMMMXASAMXSMMMMXMSMAMMXAMXMMMMMMSMMASAAXAAASXMSSMMMMAMAMXXXXXXMXAXAAXAMMSS
SMSMMMSMMSXSMMXAMMXMASMSMMSSMMMAAXAMAAXASMSAAAXSSSSMMSAMXMAAMMSMXMASASMMSASAMASXSXSXMMMSSXMASAMASXMASAMMMMSSMMAMXMAXMMSMSXAMMMMAMMSSMMSSXAAX
AXAXXXXAXMASASMSMSSXMAXAXAAAXAXXXMMSXXXXMAMAMMMXAAAMMAXMSAMSSMXXXSXMASAXMAMMSMSAMXMXXMAMAMSSSXXAMAMXMASAMAAXAXSMAMSSSMAXMAXXSAAASAMAAXMXMMSS
MMMSSSSSMMXMASAMXAXMMSXMMMXSSMMSAXXMXSMMMSMSSSXSMXMMMXXMXSMAAMXMMMMMXMMMMMMAAAMMMAMXSMSSMASAMXMSSMMXASAMMMSSSMMSMSAAASMMXSXASMSXMASMSMSXMAXM
XAXXAAAMXXXMAMAMSSXSMMASASMXAAASMMSMAXMASAAAAAAAAASAMASXXMMSMMAMXASAXXXXXSMSMSSMMASAAMAMXSMAMMXMAMMMXXMXAMXAAAMAMMMSMMXMAMMXMAMASAMXMAXMMXSX
SSSMMMMMMAXXXSXMXAMSASAMAMASMMXSXAXMSMXSXSXMSMMMMMSASMMMAXAMXMXSAAMAXSMSAXAXSAMXMAMMSMMAMXXAMAMXMAMAMXXSMSMSMSMAMMAMXSAMMSMSMAMMMXMAMAMMSAMX
XMAASXXAAMSMMMMSMMMSAMAMAMMXASAMXSMMMMMMXMXXAAAXXAXAMAASXMXSAMASXSMMMXAAMMMMMMSSMXSMAAXSMSSSSXSAMMMSMMMAAAAMXXMSMMSSMSASAAMXMASXMAMXSAXAAASM
MSSMAASXSMAAAMAXAXAMXMASASXSXMAXMAXAAAAMAMSSSSMSMSSSXXMMMAXXXMASAMAXMMSMSAXAXMAMXAXXXMMXAAAXAASMSMAAAMAMMMSMXXXXAAAAAMMMXSSXSASASMSXMMMSMXMX
AAAXMXMAAXSSMSAXSMMSSXMMAMAAAXSMMSSSSSSMAMAAMXAAAMMMSSSMMSMSMMXSMSSMMAMXSMSMSMAXSSSMXSAMMMSMMMMAAMSSSMASAMAAXMMSMMSSMMXASMMMMXSXMASMAAAAMSMS
SSMMXXMSMMXMAMMMMAAAXASMSMSXMMASXAAXAMMMSSMSMMSMXMAAAXAAAXAXAAAXXAMAMMSAMASAMXSMMAAAAMMXSAMAMAMSMMMAMMAXAMAMSAAAAMAAASMXSMAAMAMAMXSSSMMSAAAA
XAXASAMASAAMXMMAXMMSSXMAXAXAMSXMMMSMSMAAXMAAXXMXMXMMSSSMMMMMSMSSMSSSMAMAMXMXMAXAMXMMMSMASAMXMXXAMAMAMMMXMMSXMMSSMMMSXMASMXSSSXMAMXMXMAMMMXMM
SSMSMMMAMMXSXSMMSASAMXMXMAMXMSAMXXXAXMMMSMSMSMSASAAMAAAXSAMAAAAXMAAAMXSAMASMMXSAMMSSXXMASXMSXMMMSASXSMSAAMXAXMMMMAXXAXSMMMMMAMSMSAMMAAXMMSMX
XMMMSXMSMSXSASAAXMMSMSMAMSXMASAMSMMMMAMSMMAAXASASMSMMSMMSXSXMMMMMMSMMAMXMASASASAMAAMMMMMSMMMASAMSAMXMASMSMMMMSAMSMMMMSMSAAAMAMAAXAXASXSAAAXX
MXAMXXXXAXAMAMMMSMAAXAAXSAAXASXMAAAMSMMAAMMMMMMXMAMAAAMXSAMMXAMXSXXXMXXAXASAMASMMMSXMAXXXAASAMXXXSSXMMMAXXMSAMMMAAAAMAASMSSSSSSSSMMMMASMSMSM
ASASXMSMMMMMAMAAAMSSSSSMMXSMASASXSMMAXSSXMAMXSASMMMSSXSAMAMASMSXMASMMMSMSMMAMXMAASXMSMSMSSMMXXXAAXMASMMMMAXMMMMSXSMSMMXMAMXAAAXMAXAXMAMXMXMA
MASMAAAAXAAMAXSSMMAAMXAXXAXMMMXAMMMMMMAMXSXSAMXSAAAAXXMASAMXSMAAMAMAAMAAXMASXAMMMSMXAAAXAMXXSMSMSMSAMAASMSMSSSXSAMAXXSAMAMMMMMMSMSMSMXSASAMS
XMAMMSSSMMSSSMMAXMASMMSMMXSAAMMSAMSASXASXMMMXMAMXMSSMXMAAXSXMMSSMAMSXSMMSMMXSMSSMXMSMSMSMSMXMAXXAAMMMSMMAAAAXSAMAMAMASASAMXXMMAMXAAMAMMMXAMM
MXAXMXMXMMAMXASAMXMAMAXXMAXXMMAMMASASMSMMXSASMXSAMXMAMSXMMSAMXAMMAXMAXXAAAMAXXAAXXMMAAMAMAXAMSMMMSMMMXAMSMMMAMMMSMMMXMAMXMMSAMAXMMMMAMXMSMMM
SMSMMAMASMMMSMMXAXSAMXSMSMSSSMSSSMMASXMASAMASAASXMAMMMASXASAMMASMSAXMASMXSMMMMSSMSAMMMSASXSASAAAAMASAMAMXAAXXAMAMASXSMSMXMAAMSSXSAMSXMXMAXSA
XXXAXAMAMASAMASXSMMXSAMAXAAAAAAAXXSAMASXMXSAMMMSASXSASMXMMSXMXXXAXMXMMXAAMASXAAAAAMMAASMSMAMXXXMXSAMXSMSSXMSMXMASXMAMAMMXMASXMXXSAMXAMXSMSAA
MSSSMAMXSXMASAMXAAXAXAMMMMMSMMMXMXMAMXMXXMXMXXXSAMASAMXXMAMAMXMMMMXXMMMMMXAMMXSMMMXSMXMAMAAMMSSMXMASXSAAAXSAMXSASXMAMAMMASAMXAMMSAMSSMMXXAMM
XMAMSAMXAXSAMXXSSMMASXMSSSXXMXMASMSMMAMSMMASMMMMXMMMMMMAMASAMASAXMAMAAMXSMAXAAMXSXASMMMXMMASAASAMXSMAMMMMMSAMAMASXSASXMSASXMMXMASAMXXMAXAMXS
AMAMSMSSXMMMSSXAAXXXXAXAAXASMSMAMAAXSXMAAMASAAXMXXSASASASASMSASMSXASXXMASXMMMXSAAMASAASMSSXMMXXXAXXMMMXAXASMMAMXMASXMMXMAXAXSXMXSXMXSSSSXXAM
ASXMXXMMXMASAXMSMMSMSMMMSMMMAXMASMMXMASMSMASXMMSAAXASAMASASAMXMAMMMXXAMXSASMMAMAMMXMXMSAMXMSMMSSSSMMAMSMMAMASXXXSXMMSSXMSSSMSXMASAMAMMXAAMXS
MMMAMXXMAXMMAXXMXMXAAXAMAXXMMMSMSAXAXMMAMMXMMSAMMSMMMAMAMMSXXXMASAMMSXMXXMMAMXSXXMASXMMMMAAMAAXAAAXSAXAAXASMMMMXXAAAAXASAAAMXAMAXXMASXMMXSAM
MSAMXSXXXSSSXSMMXMAMMMMXMXMASAAXMASMSXAMXAAXXMMSAMXXSXMXMASAAXSASASXXAMSMXSXMMSAMSAXAXAASASXMMSMSMMSXSMSMASAAASMMXMMSSXMMSMMSMMMSMMAMAAXAMAS
XMAXASMSMAAMAMAMAMMASAMXSASMMSSSMAAAMXMXMSXXMAXMASMMMXMSXMMSMAMAXXMXSASAAMMASAMAMMMXAMSMSAMAMXXMAXXMXSAAMXMMSMAAMSSMAMAMAXAMMMSAAAMSSSMMMSMM
SMSMAXAAAMXMASAMAXXAMASASXSAAAMAXSMSMAASAMAXSAMMXMAAXAMXMSAMXXMSMXXMAMMXXASAMASXMASXMMMMMXSAMAXXMSAXAMSMMAXMXMMSMAAMMMAMMMAMAAMMMSMAAAAXAAMA
SAMMSMSMSXXSMSASMSMSSMMMSAXMMMSAMXXAXAMSAMAMSASXASXMXAXAAMASXMAAAXSSMSSSMXMMSAMAMASAAAXAMASAMASMXSAMXMAMXMMMMSSMMSSMSMSSXSAMMXSXAMMMSMXMASXX
MAMAAAASMMAMASAMXMAXAAXXMMMSSXMMSMSMASXMAMMMMASMXMASMSSSMXAMAMSMAMAAMAAAXMAXMXSMMASXMMSXMASAXAXXXMSXMSMSASASAXAAMMAAAAXAMMXMMAXMMSAAMMSMSAMX
MXMMMSMMMMAMMMXMAXXSXMSSMXAAAXXXAAMXXMASXMMAMMMMSSMMAMAMSMMSXMAMMMSSMMSMMXXMSMSMMXSXXMAMMMSMMMSMSAAXMAASASAMSSSMMMMMMXMAMMSSMMMSASMMMMMXMAAM
MASXMXXXXSASXMXSXSXSAAAASMMSMMASMSMMSSXMMXSMXSAMXASMSMASAMXAMSMSXAAAAAAXSSMMAAMASXXAXSAMXMXAMXAAMXMMSAMMMMMMAXAXAASXSXSXMMAAAXASASXSAMMAMXMS
SASAMAMSMAAMMMMMXXAMXMSSMMXAMXXXAXMXAXASAMAMAMSMSAMXXSXMASXMMAAMMMMSMMSAAXASMMMMMAMMXMXSASXSMSMSMXXAMXSASAMMXXMXMMSAMMAMXMSSMMXMAMASAMSASXSX
MASMMASASMXMAAAXMMSAMXXMASXMMSAMXMXMASMMSAAMMSMXAXSAMXXSAMAXSMMMAAAXAMXMMSXMASMXSAMSAMXSASAXXAAMAMMMSASAMMSAXMSMSAMAAAAMMMAAXXMMSMASAMSXSXAM
MMMASXSASXMSSSXSXAMAXXMXSSXMAMAXAXSAMXMAXSSMSAMMXXMAMAMSAMMAMAASXMSSSMAMXMAMAMXASAMSASMXAMAMMMXMASXXAMMXSSMXSXAAAXSXMXMXAMMXXMASXAAMMMMXMASA
MSSMMMMAMAXAXMXSMSMSMMXMMMMSSSSSXSXASAMXMAMXMASASXMAMXXSASXSMSMSAMXAAMAMASXMMSMMMSMXMXMMMMXMAXXSAMMMMXSAMAAMMMMSMMAMXASMSSSMXASXXMMSMASAMAXS
AAAXSMSMSMMSSMMSAMXMASAAAAXAAAAMXAMXMASMMXMASMMMSASXSXXSAMAXMXXMMMMSMMXMASXAMXMXAXMXSMAAXSXSSMMMXMAXXAMAXMMMXMAAXMSAMXSAAAAAASMSXMAXMAMAMXSX
MSSMXXAAAAAAAMAMAMSMAMSSSMSMMMMMMMAXXXMMAXXAAXXMSAMAMAMMAMSMMXAXMAXMAMAMSMMAMAMXMSMAMAMMXSAMXMAMMSSSMSSSMXAXXMMMSMMSSXMMMMMMMAAXXMAMMMSMMMXM
MMMMXSMSMMMSSMMXAMAMAMXMMASXMAAXASXMSSMXMMMMMSAMXMMMMAASXMMASMMMSXSSXMXSMAMMMASASAMAXMXSAMXMASASASAAXAAAMSMSAMXXMAMAMAMMAXASMMSMMXSSMAAAAAMA
XAAAAXAAMMMXMAMSSSXSXMXAMXMAMXMMXSAAAAXAAAXXSAMXAXASXXXAMXMXMAMMMAMAMSXMMMMASASMXAXASXMMXSASXSASXSMSMMSMMAMAAMSSSSMAXAMMMSMMAAXXXAMAMSMSMSAS
SSSSMMSMSAAAMXMAXAMMMSMXXAMXMASXXMMMMSMSSSXXMAMMMXMSMMSASASXSMMAMMMMMAAMAMXAMMXMSSMASXAAXSASAMAMMSXMAMAMSASXMMAAXMSSSSXSAAASMMMXMASAMMXXAXXM
AMXAMXMXSXSXSMMMXMAAAAMXXMXSXMXAAXXMXMAXAXMXMAMASAMXAXSAMASXAXSXSXAXMSMMASMAMSMXAAMMMMMMXMAMXMSMAMASMMAXSAMAASMMMMAMAAAMSSMMAAXXSASXXSAMXMXX
MMSMMAAMMXMXXMAXASXSMMSSSMASAMXSAMXMAMSMSMAMSXSAXASMMMMMMAMMMMSAMMAMMMSMASMMMAXMASMASXMMMMSMXMAMAMMMAXSMMSMSMMAXXMASMMMMXXXMSMSMMASAMMMSSMSS
XMAMMXSAMASAMSSSMSAAMSAMAMASXMAAAMXSXSXXXAXAXAMASXMMASAMXAXAMAXAMMSASAAMXSAMSXMMAMXXMASXAXAMAXMSSSSSSMMASAMXXSAMXSASXSAMXMXXMMXAAAXMASXXAAAX
MSASAAXXMAMAXAXAXMMMMMASAMASAMMSXMMAMMMAXSMSMXMASXAMMXXSXXMAMXSXMXXAMASXAMAMSAMSSSXSSMMSSSMSSSXMAAAAAASMMMSSMMAAXMASAMXMXAAMMAXSMSSXSXMSMMMS
MSAMMMASMASMMMSMMMXXXSXMXMAMAXAMAMAMAAMSXAMAAAMASMSMSMSMMSAMMAMMMAMXMXXAMMSMXAMAAXXXAAXAXAXAXMAMMMMSMMMXXXAMASXSXMSSXXAMSXSAMSMAAXXMMAMXMASA
MMSMXMAXMAMAAASMMSAMAMAXASXMSSMSSMMXXXSXMSSSSSXXXAMAAXAAASAMMXXAMSXMMXMASAMMSSMMMMSSMMMXMXMSXXSMASXXMXSXMMMSXMAMAMAMMSSMSMXXAAMMMMXMSXMAXXXM
XAMXAMASMSSSMSMAAMASASMMMMAMAAMAMXXSSMMMAAAAAAXSMMMSMSSSMMAMXAXSXMAMSASMMXSXMMASMAMAASXSXMAMXAXXAXMASAMXMAXAAAASMMAAMMXAXMMMXMMSAMXMAMSMSMAM
MSMSXMAMXAAXMAXMMXAXAXAAMSXMMMMASXMMAAXMMMXMMMMXAMAMMMMAMSMMMMSMAMXMXASXAAXAASAMMSMSMMAMMAAXMXMMMSSXMMSMSASXSMASXSSSXAMMMMXSXAMSAMXMAMAAAXAM
AAAAAMSSMMSMSMSMSMSMMSSMMSAMXASXSMAMSMMMASXSMXMMAMAMAASXMASAASAMXMMMMMMMMSSSMMXSAXMMAMXMASMMMAXAMXMXXXAMMMSAMXAXXAMXMXSASXAMASASAMXSXSMSMSXS
SMMMSMAAXAXMAAAXMAMAMAXMASAMSASAMXXMASASAMMXAAAXSSMXSXSASAXSMSXMASAXAAAXMXAMAMXMXMASMSMSAXAASASXSAMMMXMMAMMMMMXMSMSAMXSAMMASAMXMAMMMAXAXAMMX
MASXAMMMMSSMMSSSMAMAMSMSAXXMMAMAMAXSASMMMSMSSMMSXAAXMAMMMMXMXXAXMSXSSSXSMSMSAMMSAMMMAAXMASXMXASXSXSASMSSXSASASAAAAAMMAMSMXXMMSMSMXAMMMMMSAMX
SAMSAMXAXXSAAAAXMAMXMMAMASXXSMSMMMXMAMAXAAMMAAXSXMMMMAMXASMSASMMMMMXMMXMASXSASASASAMXMMMAMASMMMMMASAMAAAAMASAXMSMSMXMAMAMMAMXSAAASAASASXXSSM
MAXXMSSMMAMMMMXMMSXSAMMMAXMASAAMAXXMAMAASMSSSMMMMXSASMMSXSAMXMXAAMMXAMMMAMASAMMSAMASXSAMXSAMAXMAMXMMMSMMMMAMASAMXMMSSMMASXMMAMXMMXSMSAAXMASM
ASXMAXAAMSMSMMMAXAXMAMXMASMXMMMSSMMSASAMXAAAAMAAAXMASMAMMMXMASMSMSAMXMASMMXMSMXMXSAMXSMSXMMXMMSSMMAMXAAAXMXSXAXSXAAMASMXMAMMXXAXXMAXMAMSMAMS
SXMASMMMMAAAAASXMMSSSMMSAMXMXSMAXSAMXXXMMMMXMSSMSXMAMMAMAMAMMXAAAXMMSMAMAMXAXXXAXMMMXXSAMXSASAMXMSAMSSSMSMAMMMXSMMXSASXMSSMMASMSAAXXMAAAMAMA
XXXXMXMXSMSMSMSASAAMAAAAAMXMAAMAXMSMSMSXMXMSAXAAMXMASXSSSSSSSMSMSMXAAMASMSMSMSMSSMXMSXAXAASXSMSMMXAXXAXAMXASAXAMAXAMASAMAMAMAAXXMASXSMSAXSSM
MMMMSAMXXMXMXASAMMSMSMMMMMAMMSMMSMASAAAMSAMMXMMXMASMSXMAAAAAXAMXXMMSXXAXXAAAAXAAXMAMXXMASMMAMAAASMMMMAMMSMXMXXMMAMMMAMAMASXMMSMXXASAXAXMAMAX
AAXAXAXSMSAMMMMMMXXMMASXSMMMAAXAAMXXMMMXMASASXSMSAMXMMMMMMMMMSMSAXAXMMXMSMSMSMMMXXAXAMXMAMMAMMMSMAAXMSMXSASXSXSAXMAXMSMMMSAMSAMMSMMXMASAXSAM
SSMMXSMAAXASXAAMXASASAMAMAAMXSMSSSXMXMXXSSMMSMAAXXSAMXSAAAXAAXXXXMASMXSAXMAXMASMXMSSMSMXMAXSMMXAMSMMMAMAXAMMSASASMSMXAXAXXAXSAMAAXAMXXAAXXAA
MMMSAMMXXSSMMSSSMASAMMMSSMMXSAMXMAMSSMSXXAAAXAMMMXSAMASMSXSMSSMMXSAXAAXMXAMMMMMMSXMAXSASASXAXSMSXXASXMMMMSMXMAMAXMASXMSMMMSMSAMSSMMSSSMSMSSM
XAAMXMXSXXMAXMAAMXMMMSAMXAAXMAMXMAMAAAMMMSMMMXAMXXXAMMSXMXXXAAAAAAXMMMSXMMXSAMAAXXXXMMXMAMMMMSAMASAMXXAXMMXSMSMMMSSMMAAAAAAAMXXAAAXXAAAAXAMX
MMMSMMASMMAMMMSMMSAMAMMXSMMSMMSXSMSSMAMAAAAMASXMSSSXSXSAMXSMSSMMSMXXAAMXAAASASMMSMSSSSSMAMAXAMAMMMSMSSMSXMAXXAAAAXXAAASXMSMSMSMSSMMMSMMMMMSM
XMXAXMASAXAAAAXAASAMSMSAMXASAXAAXMAXXAMMXSSMAXAXAAXAXXMAMAAMXXXXMASXMXMSSMXSMMAMAMXAAXMMSXMMSSXMMAMSMAMAAMAMSSSMSSSXMAXAXXAAAXXAAXMAMAAXAXAA
AMSMSMAXAMMSMSSMMSSMAAMMMMXSAMMXMMSSSMMSAAAMXSSMMSMSMXSAMXSMMSSXMASAMAMXXXMSXSMSMXMMMMSMXASAXXXXMASASAMSMSSMXMAXMAXAMXSMMMSMMMMSSMMASXMSASMS
XMAXSMXMAMXAXXXAAMAMMSMXXSAMMAMXAXAAAXAMMSMMAMAAXXAAMMSXXSAAXAASMASASXSMMSXMASASAASMMMAASAMXSMMXSMSMSMMMMAMMSSSMMAMXMAMXAMAXSAMMAXMASAAAASXA
SXSMSAMSSMXMSMXMXSMMSAXXAMAMXXXMASMSSMMSXXXMASMSMMSMMAXXXMMAMXMMMASXMXSAAMAMAMAMMMMAXSMMMAMAAAAAAAXXSMXAMXSXAAMMMMSMMASASMAMSAMSMMMSMMXMXMAS
XAMAMXMAMMMMAMAMMMXXMASMSSSMSMSXMXXAXAAAMSXSAXXAAMXMMSSMMSASXSAXXMSAMAMMASASAMAMASMXMXXASXMSSSMSMMMMMAXSAMXMMXMAAXAXSMSAMMAMSAMXXAXXXSAMMMMA
MXMAMMMAXMASASASAXSSMXMAXAMAAAXAMMMMSMMMMSASXMSMXMAMAMAMASMAAXMASASXMAMSASASASXSASAMMXSMSAAXXXAXXMSSMMXMASMXAASMSSSMMXMASXMXSMMMSMSMXSAMAAAX
SASAMXMSMSXSASASASXAMMMXMAMSMXSAMXAAAAXXAMAMSAAMASAMASMMMSXMASXAMAXXMXMMASASMMMMXMAMXAXXSMMMAMMMMSAXAXSAMXXMSMSAAAXXMAMXMMSAMXAAAMAXMMSMSSSS
SASASXXMMXAMAMXMXMXSMAMXMXMAAXMAMSMSSSMAMMXMASXSASASXSXSAMXXAAMXMMMMMMMMXMXMASXSSSSMMMSMXAAMXMMAXMASMMMMMMSMMMMMMMMSSXMAMXSASMSSSSSXSAMXXAAX
MAMAAMASMMMMSSMSASAXSASAMMMXSSXMMAAAAMXSMMSMXMMMASAMAMASMXSMMMAAAASMMAXMXSXMAXAXXMXAAMAMSSMSSMSXSXAMXAAASXMAASAXAAAAAXSSSMSAMAAAMAAAMAMSMMMM
MSMSXMAMMAAAMAASAMAMSMMASAASXMASMSSMSMXMAXMXAXAMMMAMSMAMXAXASXSSSMMXSSSMASAMSMMMSMXSMSAXMXSAMXMSMMXXXSSMSASMMSAMSXSMAMMMAXXXMMMSMMMMMMXMASXX
AXAXXMAMSMMXSMMMSMSMSASAMMMMASMMAMMMAMAMSMSXXSSXSSMMXMMXMXSXMAMXXASXMASMASAMAAAAMAAXXMMXAXMXSXXAXMASMMXAMAMAAXAMMAMXXMXSMMMSMMAXXASXMXASMMAS
XMMMMMMMSAMMSXAMAXAASAMAXXMSXMXMXMASASAXMASMMMMMMAAAMXMSMXMMMSMMSAMAMAMMXSASMMMMSAMXASMMSMXASMSMSMASAAMAMASMMMSMXASAXSAMAAAAAMSXSASAXSAXXMAS
MMAAAXMASAMAXSXSMSMMMMMSMAAMASXMASASASXMMAMMAMAASMMMSASAMMAAAAXAMXSXMSSMAXMAMSAXXMSMMMSAAAXXSAASAMXSXMMMSAMMAAAMSXMMMMASXMXSSMAAXASMMMXMSMMS
MSSSSSSMSMMSMMXAXAXAAXAXMXMMXMASASMSAMMSMMMMAMSMSAMASXSASXSMSSSSXMAAXXAXMASAXMSMMXSAXAMXMSSMMXMSMSAMAMAXMASMMXMASAAMASAMAXAXAMMSMMMXXAAMAAAX
XAAAXAAXSAAMAMMMSSSSSMMSSMMSSMMMASAMXMASAMMMSXMASXMASMMAMAAAAAXAAXSXMSSMXMMMSMASMASXMXSSMAXAMSXMXSASASMSMMMMASMAMMMMMMASMMMSAMSAMSASASXSXMMS
MMSMMSMMMMMSAMXXAAAXAMXAAAXAAASMXMAMMMXXAMXMMAMXMXMASAMXMMMMMSMSAMMMASMMMAASAXMAMXSXMAMXMASMMSASASASAXAXMXMXAXMSMMSMXXAAASXSAMMAXAASMXMSMXAX
XAMAMMASMXMSAMXMMSMSMMMXSSMSSMMXXXXMMSMSMMAMSAMXAXXMSXMSXMAXXAMXAMXAXXMAASMSMMSAMMMXMASXMXSXASAMXMAMMMSMSAXMMSAMAXMASMSXXMASXMSSMMXMASAMXMSS
MXSAMSAMXAMXXMXSAMXMMAXXAXMAAXSMSMSXXAAAXXAMMAXMSMSMSAMXMSMSSMSXSMSSSMSAXXAXAASASXAASAXXAMXMMMXMAMAMXAAASMSMAMMMXMMAXAXMASXMAXAAAXSMSMSMSMMA
SMSAMMASXMSSMSXMASAMSASMMMMSMMAXAAAASMSMSSSSSXSAAAXXMASMAAAXAXMAXAAXAMMXMSXMMMSAMMSMSMMSSXMXAAAMXMAMXXMXMXXMMSXMSMMMSSMMASASXMSSMMAXAAXXMAAM
MAMXSMMMXAAMAXAMAMAMMASXMAMAMAXSMSMMXXXAMAAAAAMXMMMMMXMAMMSSXMMMMSMSMMSAMAMMXAMXMAMXSMMAMASXMSSMSSSMAMSAMSAMXSXMAAAAMAXMASXMXMMMAMMSMSMASAMX
MAMAXAXAMMMMMSXMAMAMXAMXSAXAMXXAXAXSSMSXXMMMMMMXXMASXMASXXAAAMXAXXAMXASASAAMMSSMXXSMSAMXSXMMMMAMAAAMAMSASAXSMXMSSSMXSASMXSXMAXMXSXXAAAXAMXMA
SMSSSMMMSAMSASMSMSSSMXMXSASMSXMASMSAAAMXASXMXSAMXMAXAMAMAMMSMMSSSMAMMMMAMXSXAMAMXMSAXMMMMAAAASAMMSMSMXSXMMXXAAAXAMXASAMMASASAXSAMMXMXMMASXAA
AAAAAXAXMASMASAAMAMXAASAMXMAAMAXAMMMMAMMXMAMAMASMMSSSMASAMXAXXXMAXAMAMSMSAXMXSAMAMMXMMAAXSSSMSXMXAMMMAXAAXAMMMXMASMMMAMMASMMASMASAASXSMXSAMX
MMMSMMAXMSMMXMXSMXSXSXXASXMMMMMMSXMXSAMXASXMAMAMAAXAASASXSXSSXMASMMSAMASAAXAASMSXMAMSSSSXMAXMXAMSMSAMXSSMXMSAAXSAMAXSXMMASAMXMXSMMXSAAXAXXXX
XAXAMASXMXAMXSAXMAMMMMAMXAXMASXAASMASAMSAMMASMMSMMMMMMMSXAAMAMXSAMXSXSXSMXMMMXAMMMMMMAXMXMMXMMXMAASAMAAAMSASMSMMASXMAMAMASMSSSXMAXXMXMMMSMSM
SSSMSAMAASAMXMASMASAAMAMSMMSASMSMXMASAMXAMSAXAMXXMXXMSXMMMXMXAMXAAMSASMMXMSXSMAMAASMMMMSXXXAAXMMMXMSMMMAMSAMAAXSXMMSAAXMAMXAAMAMAMAMXXAXAAAM
XAAXMASMMSAMAMSMMASXSMMXAAMMAMAMMSMASMMSAMMASXMMSMSAXSAXSMSXSAMSXMXMAMMAAAMAMSAMMMSAAMAMAASMMMXAXAAXXMSSXMAMSMXAMXAAASXMMMMMSMASXSXSXMSSMSMS
AMMMSXMAASXSAXMAMXSAMXSXSMMMSMAMAAMXSXMAMXSAMASMSAAMMSAMAAAMXAMXXSSMAMSMMSMAMMXSMMSXMMAMSMXMAXXASMSMAMAMASXMMMMXAMSMXMMMSASAMMXSAAXAAMMMAAAA
XAXXSXXMASAMMSMMMAMXSASAMXSAXXXMSMSAMXSASAMXSXMAMXMXAMXMMMMMSMMXAAASAXAASXMMSMMSAAMSSSMXXXXSASAMXAMMAMAMXMMAAAMMXAMXAMAAMMMAMMMMMMMXAMAMMMSM
SXMAMASXAMAMAXAAMASMMAMAMAMASXMMXXMXSASMMXXMXMMMSXXMSSSSXMSAAAXMMSXMSMSXMAXAAAASMMMAMAMXMMXMASMMMSMSXXMXSAMSSXMAASASASMSSXMAMXXASXSXXSASAAXX
AMMAMMMMMSMMXSSMSXSXMAMAMAMAMMSAMXSMMMSAASMSMAAAMXMXAAXXMASMSMMXAAAXXXMASAMXSMMMXAMXSMSMAXAMAMASAMAXMXXXSAMAAMMSXMAMAAAMAMSSSMSASASAMSASMSMA
XXSAMXAXMAASMXMASXSMSXSASXSASXXASXXXAAMMXXAASMMSSMAMMSMMMAMAXASMSSSMMXSXMAXXAAXMMXSAMXXXSMMMMSAMASMXMAXXSAMMSMAXAMXMAMAMAMMAAAAAMXMAMMMMMAAM
SXSASXSSSMXSSMSASMSAMASAXMAMXXSAMXMXMXSAAMSMMMSAXASMAMAXASMAMASAMAAXSAMAMASMXMMSAAMASMSMMXSAMMXSAMMMAMXASXMAMMASXMXXAXASASXMMMMSMMXSMSASMMSS
MASXMXAMAMXSAMMASAMAMXMAMSSSMXMASASAAASXSXMAXMMAMSAMXSXMAMMMMAMAMSMMMASXMASAASAMMSMSMMAXAAXASMAMXSXMASAAMXXAXMASXMMMASASASMMMAXMASMMAMMMAMXA
MAMAAMXSXMXMAMMXMMSMMSMXMAMAMAMMSASMSMMAMXSAMXMXMMMSASASXXSAMXSXMAASMMMXMASMAMASAMXMASMMMSXMMXMMMAAXAMMMMXMSSMASMSAAMMMMMMAXSSMSAMAMAMXSAMSM
MASMMMAMMSMSSMMSMXMXAAXXMXSASXSAMXMXMAMAXXMAMXMXMXXMASAMXMSXMAXAXXMMAAAXSAMXXSXMAMASAMXMXAASASXAXSSMSSXAAXXAXMMMXMXXXAXXAMSMMMMMAMXXAMXMAMSX
MMAMAMAXAXAXAAXMXAMMMSSSMAMASXMASAXAXXMMSMMASAMAXMSMMMMMXXXSMSSMMSSSSMMMXXSXXMMMSMMMAMXXAXXMAMSAMXAAAAXAMXSMSMSSMMSXSAMMMMMAAMXSXMSASMSSXMMM
SAAMMSAMSMSASMMMSMSAAMAMMAMMMMXAMXSMSSXMAMAAXAMASXSAMAXMXAAXMASAAAAXXASXMMMMMAAAAAXSAMMMSSXMAMAMASMMMSMMSMAAMMAAAAXAMXMASASAMXAMXMSAMAMAMSAS
XMSMXMAMAAXAXXAMAXMMSMAMSXSSXSSSSXSXAMASXMMSSSMSSXXAMXXMSMSMMASMMMSMSMMAAXAAMAMMSMMMAXMAAAASMSAMXMXAAXMXAXMAMMXSMSMMMAMMMASMXMXSAXMXMXMSMSAS
MAAMXSMMMMMAMSSSSSXAAMMMMMAMXMAMAMXMXSXMSAAXAAMXMASMMSXAAXXAMAMAAAMXSASXMMSXSAMXXAMSSMMMSXMMASXSAMMMSSXSMSXSSXMXAXAASASAMMMXAASXXSXXAXAAAMAM
SMMXAMXAAAMAMAMAMAMXSSMAXSAMAMMMXMXSMMMASMMMMMMAMXMAAMMMMAMXMASMMXSAXXMAXXAXSASASAMAAAMAMAXMXMMSASAAXXAMXMMAAXAMSMSMSASMSMSMXSAXMXMMSSSMSMAM
AAXMASMSXSSMMSMMMAAXMAMAMSASXSAMASXMAAMXMAAXXAMMXXXMAMAXXMAMMXXAASMMMSSSMSMMSAMASMMSSMMASAMXAXXSASMXMMXMMXMAMMSXMAMXMMMAXAXAXMASMAMAAAAAMXMX
SAMAASAXAAAXSAMMSXSASAMXASXMMSASAXMSSMXMXSMSSXSAASMSSSXSXMSMMSSMMXASXXAXAAXAMXMXMAAAXAMAMAMSMSMMXMASXMAMSXMAXAMMMXMXMASMMMMSXMSXMASAMMMMSASM
AXMSXMAMMSMMSMSAXMMAASMMMXASASXMMSXAAAAXXMXAAAMMMMAAAMASASXSAAAXASMMMMAMSMSXAXSMSMMSSXMMMMMAXXMASASAMXAXAXMAMAXAXMASXMSXMXAXXMXMSMMAXXXXXAMX
MMMMXMAMXXXAXXMXMASAMXXXXXXMASMSMSMXSXMSAMMMSMXASMXMXMASMMAMMSMMMXXAAMMMAXMMSXMASXAXMAXMAXSMMMXXAXXAAMMSMSSSSSSMSMAMAXSAMAXSASXMMXSAMSXMMMMX
XAAMMSXSAMMSSMMSSXMAXMMMSMSMXMAAASMMMAMSAMAAAXSXSMSXXMAXXMMMAAMAXASMSSMMMMXMMAMAMSAXSXMXXXAXXSSSMSMMMMXAXMAAAAAMMXSSMMSAMAXSXMASAMMSMSAMMSMM
SSSXAAMMASAAAAAXMASXMMAAMAXSSMSMSMMAMXMXMASMMXAAMASAMMSSXXAMSSXXMASMAMAMXMAMMSMSXSXMXXSAXSSMAXAAXAAAAASAMMMMMSMMAAMAMXSAMAXXXXXMASAXMXAMAAXA
XMMMSSMSAMMSSMMSMAMMAMSSMSMAMAAMAAMMSAMAXXAMXMMMMAMSMAXXMMAXMMMSAMXMAMAMAMAXAAXXAMMMSAMXMMAMMMMMSSSMSXSASAXXAXXMMMXXMXSAMXMMMSMSMMMSMSSMSSSS
SXAAAAAMAMXAAAXSMXMXMMAMXXAAMMMSSSMASASAMSAMXXASMMSXMXSAMXSXSAAXMXXSXSSSSSSMSSXMAMASMXMSASASAAAAXAAAXAXASAXMMSXSAXSMAASXMAAASAMSASMMMAAXAAAX
ASMMSSMMASMSMSMMXXSASXMSMXSSSXXMMMMXSAMXXMSMSXXMAMXMAMMAMXXAMMSSXMMMMSAAAAXMAXXSAMXXAAMXASASXXMMSSXMMMMMMMSMAAAMXSAAMASMXMMXSASMXMAAMSSMMMMM
MXSAXAAMXSAMSMAMAAMAMASAXXXAMMAMAAMXSAMXXAAXAMXSSMSMSSSSMXMXMAXMAXAASMMMMMMMSMMMAMSSSMSMMMMMMXXXAMAXSXMXAAAMMMSAXMASASMMSSMMSMMXXSXMMAXAMXMA
SAMMMSMMXMMMASAMSSMAMXSASMMAAAASMXSASAMMMSMSMSAAXSXAXAAMAMSMMSSSMSMMXAMXMAMXAMSSSMAAMMXXAAMAMXSMAXXMSAMSMSMSXAMMMMMMXAAAAAAMMXASMSXXMAXMMAXS
MAMSAMXSASASMSSXMMMMSMMAAXSSMSASAAMASAMXAAAXAMMSMMMAMMMMAMAAAAAAAAXSSSMASMSAMXAAXAMXMASXSMMASAXSSMSMSAMSXMAMMMSAAAAAMMMMSSMMAMXSASMSMSMSMSSM
XAMMASAMASASXXMXAASMAXMMMXAMAAAMMMMAMAMMSSSMMMXAAXASMAXSASXSMMSMSMMMAAMXXAMMXMMSMXSAMXMMAXSXMMMMMMMASMMMASXMAASXSSMSAXAAXMAMXXMMXMAAAXAMAXAM
XSXMAMMMMMAMAXMXMMSSXSXAXSAMXMXSASMXSXXAMAMAXMMMSSSMSSXSXSAMAAMAAXAMXXMMMSMMAAMMAASXSAXMSMSASXAMAAMMMSSMMMAMMXSAXAXAXMMMSMSXSAMXMSSMSMAMMMAM
MMAMSSXXAAAMXMAXSAMMMSMSMSAMXMAAMMAASAMXSMSXMMXSAMXXAXAMAMXSAMMSMSXMAMMXMMASMSMMMMMXSMSSMAXMSMMSSSXSAAXMMSXMXAMXMXSMSSMXSAMAMXMMXAAXXXAMXSSM
ASAMXMASXSMSXASAMSXXAMXAASMMAMMAMSMMSXSAMXSAAAMMASXMXMMMSMAMAMAXAMXMASMAASAMMAXAXAAAXXXMMXMXSAXAAMAMMMSXMASMMSSSSMSXAAMAMAMMMSMMMXSMSMSMAXAM
MSASMSMXAXXXXSAXMASMSSSMMMASXSXAMXAXSAMXSASMMMXSAMAXAAMAXMXSAMXMAXAMAXMAMMMSSMSMSMMMSMMMSMSAMXMMXMAMMMMMMASMAAAXMAMXXSMASAXMASAXSMMASAMMSSMA
ASAMAMMMSMSSMMAMXAXXAXAMAMASAMXSMSMMMMSAMXXAMXXMASMMMSMSSMASASXMSMSXSSMXSAAXAMXAAXAXAAAXAAMAXSMMSSMMSAAAMAMXMMSMMSXMAXMMSMSMXSSMMASAMXMAAXXX
MMMMMXAMAAXXAXMMMSSMMXMXASXMAMAMASXMXAMMSXSSMSMXMXXMAAXMAMMSAMMAXAMAXMAASMXSAMMSMSMSSSMSMSMSMSAAAXXAMXMSXMSAMXXXAAXAMXAMXMXMASAXMAMAMMMMSSSM
XAXSSSSSMSMMSMMAXMAMMASMMSMSAMMSAMAXMMMSAAAAAXAASMMMSSSSSMXMMMMAMAMSMMMMSAMSAMAXMXMAXAAXAXAMASMMSSMSSSMMAASXMASMSSMSXSSMMMMMMMMSMMSXMAAAXAXA
MSXXASXAAMMAAASAXSAMSASAMXASMSXMASXMSXMSMXMMMMAMSAAAAMAAAXSAMXMXSAMXASMAMAMSAMSMSAMXSMMMMMAMAMAAMXAAAAAMMMMAAASAMXAAMAMXAAAAAAXAAASASMSSMAMM
MXAMXMSMMMMSSMMSAMXXMMSXMAXXMAMSAMXXSAMXXSXASXMXSXMSSXMSMMSMMAXMAMASAMXSSSMXMAXASXSMXXXXSSXMMSMMSMSMSMMMAASMMSMXMMSMXMASMSSSSSSSSMSXMAXXMAMX`