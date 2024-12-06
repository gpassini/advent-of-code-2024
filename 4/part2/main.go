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
	if m[y][x] != 'A' {
		return 0
	}

	masCount := 0
	var axises []axis
	for _, dir := range crossAxises {
		if tryDir(m, x, y, dir) {
			// fmt.Println("Found MAS around (", x, ",", y, ")", dir)
			masCount++
			axises = append(axises, dir)
		}
	}

	crossCount := 0
	switch masCount {
	case 2:
		// 2 axises might not form a cross
		if axises[0].makesCross(axises[1]) {
			crossCount = 1
		}
	case 3:
		// 3 axises are sure to form one cross
		crossCount = 1
	case 4:
		// 4 axises are sure to form 2 crosses
		crossCount = 2
	}

	// fmt.Println("Crosses:", crossCount)
	return crossCount
}

func tryDir(m lettersMap, x, y int, dir axis) bool {
	firstDir := dir[0]
	if r, _, _, ok := m.move(x, y, firstDir); !ok {
		// out-of-bounds
		return false
	} else {
		switch r {
		case 'M':
			oppositeDir := dir[1]
			r, _, _, ok := m.move(x, y, oppositeDir)
			return ok && r == 'S'
		case 'S':
			oppositeDir := dir[1]
			r, _, _, ok := m.move(x, y, oppositeDir)
			return ok && r == 'M'
		default:
			// can't form a "MAS"
			return false
		}
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
	case northEast:
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
	case northEast:
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
	northEast
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
		northEast,
		northWest,
		southEst,
		southWest,
	}
)

type axis []direction

func (a axis) makesCross(other axis) bool {
	return (a[0] == north && other[0] == east) ||
		(a[0] == east && other[0] == north) ||
		(a[0] == northEast && other[0] == northWest) ||
		(a[0] == northWest && other[0] == northEast)
}

var (
	northEastSouthWest = []direction{northEast, southWest}
	northWestSouthEast = []direction{northWest, southEst}

	crossAxises = []axis{
		northEastSouthWest,
		northWestSouthEast,
	}
)

// expected: 9
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