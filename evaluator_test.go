package scheme

import (
	"math/big"
	"testing"

	"github.com/Algemetric/HERatio/Implementation/Golang/laurent"
	"github.com/Algemetric/HERatio/Implementation/Golang/oracle"
	"github.com/Algemetric/HERatio/Implementation/Golang/params"
	"github.com/Algemetric/HERatio/Implementation/Golang/sim2d"
)

func TestBFVSAdd(t *testing.T) {
	// Create parameters.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		t.Error(err)
	}
	// SIM2D codec.
	sc, err := sim2d.New(p)
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := Setup("PLBFV32.kc", o, p)
	if err != nil {
		t.Error(err)
	}
	// Cipher.
	cip, err := NewCipher(kc)
	if err != nil {
		t.Error(err)
	}
	// Evaluator.
	eval := NewEvaluator(kc)
	// Case: message (12345.678) + scalar (42.122).
	m := 12345.678
	s := 42.122
	// SIM2D encode.
	mpb := sc.Enc(m)
	spb := sc.Enc(s)
	// Encrypt message.
	c, err := cip.Enc(mpb)
	if err != nil {
		t.Error(err)
	}
	// Addition.
	cr, err := eval.SAdd(c, spb)
	if err != nil {
		t.Error(err)
	}
	// Decrypt.
	crd, err := cip.Dec(cr)
	if err != nil {
		t.Error(err)
	}
	// SIM2D decode.
	mrd, err := sc.Dec(crd)
	if err != nil {
		t.Error(err)
	}
	// Check result.
	if mr := m + s; mrd != mr {
		t.Errorf("expected %f for %f + %f, but got %f", mr, m, s, mrd)
	}
}

func TestHERatioSAdd(t *testing.T) {
	// Case: message 0 (12345.678) + AdditiveScalar (42.122).
	// Parameters.
	p, err := params.New(params.PLHERatio16)
	if err != nil {
		t.Error(err)
	}
	// Keychain.
	kc, err := Setup("PLHERatio16.kc", new(oracle.Oracle), p)
	if err != nil {
		t.Error(err)
	}
	// Evaluator.
	eval := NewEvaluator(kc)
	// Additive scalar (42.122).
	as := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 1, 2, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	asb := make([]*big.Int, len(as))
	for i := 0; i < len(asb); i++ {
		asb[i] = big.NewInt(as[i])
	}
	// Ciphertext 0.
	c0p := [][]int64{{3928319131, 2322583730, 3128333473, -3588018546, -3023374777, -4583068212, 793865758, 2493685852, 4356875931, -1446151930, -155195535, 836686378, -2066582784, -4566179797, 2012557833, -4775884647, -2936070142, 2890040606, -3088406786, 1815033680, -2182161000, -2186917847, 35413812, 3547097307, 2678934739, -370357249, 3991359443, 2430971629, 454123781, -3417055595, 4371363907, 704824141},
		{-1556256741, -1792426624, 4406608444, -616705210, 2610504891, -1343850320, 1290341989, 3158043354, -3566746457, -3750192107, -4155437196, -1755722380, 1735554954, -2781807856, 2360617235, -3849070249, -3195929512, 2081665431, 178354185, 2509028358, 983465813, 1897902712, 1876620253, -2041701257, 4906960634, -642398707, 610816366, -3631039280, -2786954504, 636686208, 1293593149, -2004632067}}
	c0pb := make([][]*big.Int, len(c0p))
	for i := 0; i < len(c0p); i++ {
		c0pb[i] = make([]*big.Int, len(c0p[i]))
		for j := 0; j < len(c0p[i]); j++ {
			c0pb[i][j] = big.NewInt(c0p[i][j])
		}
	}
	// Scalar addition.
	cr, err := eval.SAdd(c0pb, asb)
	if err != nil {
		t.Error(err)
	}
	// Check result (12345.678 + 42.122 = 12,387.8).
	sum := [][]int64{{3928319131, 2322583730, 3128333473, -3588018546, -3023374777, -4583068212, 793865758, 2493685852, 4356875931, -1446151930, -155195535, 836686378, -2066582784, -4556910419, 2021827211, -4771249958, -2926800764, 2908579362, -3088406786, 1815033680, -2182161000, -2186917847, 35413812, 3547097307, 2678934739, -370357249, 3991359443, 2430971629, 454123781, -3417055595, 4371363907, 704824141},
		{-1556256741, -1792426624, 4406608444, -616705210, 2610504891, -1343850320, 1290341989, 3158043354, -3566746457, -3750192107, -4155437196, -1755722380, 1735554954, -2781807856, 2360617235, -3849070249, -3195929512, 2081665431, 178354185, 2509028358, 983465813, 1897902712, 1876620253, -2041701257, 4906960634, -642398707, 610816366, -3631039280, -2786954504, 636686208, 1293593149, -2004632067}}
	for i := 0; i < len(sum); i++ {
		for j := 0; j < len(sum[i]); j++ {
			if sumB := big.NewInt(sum[i][j]); sumB.Cmp(cr[i][j]) != 0 {
				t.Errorf("expected %s for position [%d][%d] but got %s", sumB.String(), i, j, cr[i][j].String())
				break
			}
		}
	}
}

func TestBFVAdd(t *testing.T) {
	// Create parameters.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		t.Error(err)
	}
	// SIM2D codec.
	sc, err := sim2d.New(p)
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := Setup("PLBFV32.kc", o, p)
	if err != nil {
		t.Error(err)
	}
	// Cipher.
	cip, err := NewCipher(kc)
	if err != nil {
		t.Error(err)
	}
	// Evaluator.
	eval := NewEvaluator(kc)
	// SIM2D encode.
	// Case: Message0 (12345.678) + Message1 (947.1273).
	m0 := params.M0
	m1 := params.M1
	m0pb := sc.Enc(m0)
	m1pb := sc.Enc(m1)
	// Encrypt messages.
	c0, err := cip.Enc(m0pb)
	if err != nil {
		t.Error(err)
	}
	c1, err := cip.Enc(m1pb)
	if err != nil {
		t.Error(err)
	}
	// Addition.
	cr := eval.Add(c0, c1)
	// Decrypt.
	crd, err := cip.Dec(cr)
	if err != nil {
		t.Error(err)
	}
	// SIM2D decode.
	mrd, err := sc.Dec(crd)
	if err != nil {
		t.Error(err)
	}
	// Check result.
	if mr := m0 + m1; mrd != mr {
		t.Errorf("expected %f for %f + %f, but got %f", mr, m0, m1, mrd)
	}
}

func TestHERatioAdd(t *testing.T) {
	// Case: message 0 (12345.678) + message 1 (947.1273) = 13,292.8053.
	// Parameters.
	p, err := params.New(params.PLHERatio16)
	if err != nil {
		t.Error(err)
	}
	// Keychain.
	kc, err := Setup("PLHERatio16.kc", new(oracle.Oracle), p)
	if err != nil {
		t.Error(err)
	}
	// Evaluator.
	eval := NewEvaluator(kc)
	// Laurent codec.
	lc := laurent.New(p)
	// Cipher.
	cip, err := NewCipher(kc)
	if err != nil {
		t.Error(err)
	}
	// Message 0 (12345.678).
	m0 := lc.Enc(params.M0)
	// Message 1 (947.1273).
	m1 := lc.Enc(params.M1)
	// Ciphertext 0.
	c0, err := cip.Enc(m0)
	if err != nil {
		t.Error(err)
	}
	// Ciphertext 1.
	c1, err := cip.Enc(m1)
	if err != nil {
		t.Error(err)
	}
	// Ciphertext result for addition.
	cr := eval.Add(c0, c1)
	if err != nil {
		t.Error(err)
	}
	// Decrypt result.
	mrb, err := cip.Dec(cr)
	if err != nil {
		t.Error(err)
	}
	// Decode result.
	r := lc.Dec(mrb)
	// Check result (12345.678 + 947.1273 = 13,292.8053).
	// TODO: fix variables.
	er := 13292.8053
	if r != er {
		t.Errorf("expected %f for %f + %f, but got %f", er, 12345.678, 947.1273, r)
	}
}

func TestBFVSMult(t *testing.T) {
	// Create parameters.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		t.Error(err)
	}
	// SIM2D codec.
	sc, err := sim2d.New(p)
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := Setup("PLBFV32.kc", o, p)
	if err != nil {
		t.Error(err)
	}
	// Cipher.
	cip, err := NewCipher(kc)
	if err != nil {
		t.Error(err)
	}
	// Evaluator.
	eval := NewEvaluator(kc)
	// Case: message (12345.678) x scalar (4).
	m := params.M0
	s := int64(params.MS)
	sb := big.NewInt(s)
	// SIM2D encode.
	mpb := sc.Enc(m)
	// Encrypt message.
	c, err := cip.Enc(mpb)
	if err != nil {
		t.Error(err)
	}
	// Multiplication.
	cr := eval.SMult(c, sb)
	// Decrypt.
	crd, err := cip.Dec(cr)
	if err != nil {
		t.Error(err)
	}
	// SIM2D decode.
	mrd, err := sc.Dec(crd)
	if err != nil {
		t.Error(err)
	}
	// Check result.
	sf := float64(s)
	if mr := m * sf; mrd != mr {
		t.Errorf("expected %f for %f x %f, but got %f", mr, m, sf, mrd)
	}
}

func TestHERatioSMult(t *testing.T) {
	// Case: message 0 (12345.678) x multiplicative scalar (4).
	// Parameters.
	p, err := params.New(params.PLHERatio16)
	if err != nil {
		t.Error(err)
	}
	// Keychain.
	kc, err := Setup("PLHERatio16.kc", new(oracle.Oracle), p)
	if err != nil {
		t.Error(err)
	}
	// Evaluator.
	eval := NewEvaluator(kc)
	// Laurent codec.
	lc := laurent.New(p)
	// Cipher.
	cip, err := NewCipher(kc)
	if err != nil {
		t.Error(err)
	}
	// Multiplicative scalar (4).
	ms := big.NewInt(params.MS)
	// Message 0 (12345.678).
	m0 := lc.Enc(params.M0)
	// Ciphertext 0.
	c0, err := cip.Enc(m0)
	// Scalar multiplication.
	cr := eval.SMult(c0, ms)
	// Decrypt.
	mrb, err := cip.Dec(cr)
	if err != nil {
		t.Error(err)
	}
	// Decode.
	r := lc.Dec(mrb)
	// Check result.
	// TODO: check variables.
	if (params.M0 * params.MS) != r {
		t.Errorf("wrong scalar multiplication result for HERatio.")
	}
}

func TestBFVMult(t *testing.T) {
	// Create parameters.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		t.Error(err)
	}
	// SIM2D codec.
	sc, err := sim2d.New(p)
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := Setup("PLBFV32.kc", o, p)
	if err != nil {
		t.Error(err)
	}
	// Cipher.
	cip, err := NewCipher(kc)
	if err != nil {
		t.Error(err)
	}
	// Evaluator.
	eval := NewEvaluator(kc)
	// Case: Message0 (12345.678) x Message1 (947.1273).
	m0 := params.M0
	m1 := params.M1
	// SIM2D encode.
	m0pb := sc.Enc(m0)
	m1pb := sc.Enc(m1)
	// Encrypt messages.
	c0, err := cip.Enc(m0pb)
	if err != nil {
		t.Error(err)
	}
	c1, err := cip.Enc(m1pb)
	if err != nil {
		t.Error(err)
	}
	// Multiplication.
	cr, err := eval.Mult(c0, c1)
	if err != nil {
		t.Error(err)
	}
	// Decrypt.
	crd, err := cip.Dec(cr)
	if err != nil {
		t.Error(err)
	}
	// SIM2D decode.
	mrd, err := sc.Dec(crd)
	if err != nil {
		t.Error(err)
	}
	// Check result.
	if mr := m0 * m1; mrd != mr {
		t.Errorf("expected %f for %f x %f, but got %f", mr, m0, m1, mrd)
	}
}

// TestBFVMult2048 tests the multiplication of 2 ciphertexts with secure parameters.
func TestBFVMult2048(t *testing.T) {
	// Create parameters.
	p, err := params.New(params.PLBFV2048)
	if err != nil {
		t.Error(err)
	}
	// SIM2D codec.
	sc, err := sim2d.New(p)
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := Setup("PLBFV2048.kc", o, p)
	if err != nil {
		t.Error(err)
	}
	// Cipher.
	cip, err := NewCipher(kc)
	if err != nil {
		t.Error(err)
	}
	// Evaluator.
	eval := NewEvaluator(kc)
	// Case: Messages.
	m0 := params.M2
	m1 := params.M3
	// SIM2D encode.
	m0pb := sc.Enc(m0)
	m1pb := sc.Enc(m1)
	// Encrypt messages.
	c0, err := cip.Enc(m0pb)
	if err != nil {
		t.Error(err)
	}
	c1, err := cip.Enc(m1pb)
	if err != nil {
		t.Error(err)
	}
	// Multiplication.
	cr, err := eval.Mult(c0, c1)
	if err != nil {
		t.Error(err)
	}
	// Decrypt.
	crd, err := cip.Dec(cr)
	if err != nil {
		t.Error(err)
	}
	// SIM2D decode.
	mrd, err := sc.Dec(crd)
	if err != nil {
		t.Error(err)
	}
	// Check result.
	if mr := m0 * m1; mrd != mr {
		t.Errorf("expected %f for %f x %f, but got %f", mr, m0, m1, mrd)
	}
}

func TestHERatioMult(t *testing.T) {
	// Case: message 0 (12345.678) + message 1 (947.1273) = 11,692,928.6708094.
	// Parameters.
	p, err := params.New(params.PLHERatio16)
	if err != nil {
		t.Error(err)
	}
	// Keychain.
	kc, err := Setup("PLHERatio16.kc", new(oracle.Oracle), p)
	if err != nil {
		t.Error(err)
	}
	// Secret key preset big.
	skp := []int64{1, 0, 0, 0, -1, 1, -1, 0, -1, 1, 1, -1, -1, 1, 0, -1, 1, 1, 1, 0, 0, 1, 0, -1, 0, -1, 1, 0, 0, -1, 1, 0}
	skpb := make([]*big.Int, len(skp))
	for i := 0; i < len(skp); i++ {
		skpb[i] = big.NewInt(skp[i])
	}
	kc.SK = skpb
	// Public key.
	pkp := [][]int64{{-4216387067, 1836069869, 2001516335, -3340617436, -1836444424, -2329243897, -3398349639, -2389753658, -538912150, 912064576, 3963899003, -4548007374, 4387623173, 560753886, 3678115043, -4662116239, 4037345779, 3864796148, 3246898652, -1681759155, -1408805783, 1107088450, 939512626, -2391561032, -2018740806, 4359243558, 1210139051, -1799894561, -628280241, -1846844096, 4350890023, -1553927766},
		{3047662976, 3648351800, -4723214805, -589412704, -2752366317, -4396812798, -3822959355, -4533739225, 1223700754, -4723897659, -3480506128, 3340699993, 1866454384, -2874149000, -1957459557, 1008485621, 2463063424, 1841973385, 2063260407, -586617775, -3878219761, 3602308653, 3609153671, -716226391, -1465250363, -2004232445, -4336824147, 2790860168, -3274517433, -2236599955, -836182508, -1477538473}}
	pkpb := make([][]*big.Int, len(pkp))
	for i := 0; i < len(pkp); i++ {
		pkpb[i] = make([]*big.Int, len(pkp[i]))
		for j := 0; j < len(pkp[i]); j++ {
			pkpb[i][j] = big.NewInt(pkp[i][j])
		}
	}
	kc.PK = pkpb
	// Evaluation key.
	ekp := [][][]int64{{{-592378270, 1232887901, -4240076597, -2507950966, -244036067, 4266860198, -3795647498, -387829448, -1063435371, 2716722646, 232312952, 2616751081, 3829735507, 3605375054, 1241720635, 1727754468, 2933651026, 4669113929, -3178751746, -1185439131, -373802894, -3506534634, 3392101160, 4431994126, -3499469738, -111633455, -1021062974, -703800300, 1206085201, 3896510320, 4232777591, -4232952828},
		{1720007324, 4174361130, -1044965507, 3924438480, -3456480106, -2973544064, -3975934511, 1890543970, 4669003711, -3418588942, -3513411606, 3781687093, -2976239293, -4536487024, -1998071673, 3926019078, -1106055272, -1820219249, -3194291465, -4463946424, 4251051716, -1009373237, -275264562, -1827095842, 1829765302, -4295061476, 4378397377, -2306047404, 1651000318, 3865841105, 1836734621, 3071384110}},
		{{-965898273, -871535145, -2829273207, 3810207485, -1404723649, -2603336420, -1867004204, 4310464815, 1376129458, -4816611426, 4666530454, -2857092814, -1186824247, 2730258733, -3209970251, -4481497920, -3445169105, -2782093163, 2237478375, 4204085441, -3100154820, -589258349, 2825152665, -1474803129, -24431576, 1725089938, -4568770355, 4749657532, 2092653116, 4819319917, -4830571137, 2764063555},
			{-2774622614, -497553969, 22441196, -3103019905, -3390037273, 229969976, 2709838917, -2112035363, -2039992369, 4202390933, -2453635591, -1807796033, 231785954, 4498367660, 2131104630, -3591973335, 1330156222, -1131669242, -3735087423, -1056635912, 3412051809, 2896043637, -3594075230, 3002922704, -1026817562, 3177548642, -4239652937, 83285209, -338568473, 3444244013, -37284937, 3726560895}},
		{{688057481, -4899613482, -2491780828, 227033383, 3832543509, 1057634936, 3420248410, 4609943255, -2660856166, -1892897142, 4123503596, 92890820, 864913439, -1505737618, 3173641006, 1678019941, -1059840650, -206323968, -3791630476, 1202064711, -2662962743, 4935430210, 4113317482, 396190326, 4167189228, 831989832, 3097502995, 942178133, -1118773818, -3230376297, -2697636221, 2438670837},
			{-1763011266, 1047691812, -4087577526, -1396648061, 2290452197, 4753853664, -192274155, -1714872255, -4766350798, 3396078051, 3631661698, -1285158152, 2335783894, 4374869900, -91625455, 2882905604, -532196749, 4770233642, -2305309808, -27662999, 2221294577, -2598245449, -1433084070, -3964915712, -3114653801, 1853458215, 4276493177, 3833978274, -157465609, -431110507, 4155495448, -2126299762}},
		{{-2199048447, 1214930405, 702160830, 2645668784, -4324265180, -710444988, 4045688200, 4262805886, -4853244875, 3564411811, 2947983618, -2602216704, -3158794618, 94131916, 4795511048, -4662762378, 4624142681, -2331476963, -4123314628, -1777242813, -2145941020, 4700412928, 2901516844, -1156574387, -3690843919, 3444482550, 4921937841, -3771032622, 4434937623, 3848115841, 1311907729, -1292651793},
			{4128003758, -198265992, -3008314755, 94883595, -3239724824, -2045063722, 2123094016, 4613238440, 2129698390, 4465843979, -16409403, 4785295300, 510486989, -3034271733, -907661611, -709179980, 2831665608, -3001470646, 3572360263, 967937777, 3955336576, 4337825117, 4000180017, -2997610194, -418228098, -2976031888, 3593358232, -3334548540, -3500489897, 4648595817, 4344587316, -4222104596}},
		{{-822581333, -11711251, -4343916134, 3886361627, 616446636, -3022649822, -966271647, 3970329745, 4540970965, -2838860294, 1291466436, -1410166267, 1263571796, -575437479, 1250946354, 3583779061, 1755918059, -2163124361, -3514670672, 2201005619, -466405191, 1156315712, 2729443927, 1261292201, -339374609, -4198210378, 1598871678, 3613671203, -2789632307, -4060088784, -330200372, 4198838200},
			{2914257978, 4144502886, 1777475312, -1175333834, 214499981, -1480435705, -1190914867, 828421202, 2897100994, -1815234520, 4767883567, -1842680856, -1566356417, 1090057226, -3518759459, 11729732, -3662566075, 3856317916, -769627540, -4180871204, -2295425101, 4294787420, 3307394379, 1234628285, -3555324053, 2625906568, 359020980, 4286938445, -3285712587, 1963962101, 1758771712, -2274053020}}}
	ekpb := make([][][]*big.Int, len(ekp))
	for i := 0; i < len(ekp); i++ {
		ekpb[i] = make([][]*big.Int, len(ekp[i]))
		for j := 0; j < len(ekp[i]); j++ {
			ekpb[i][j] = make([]*big.Int, len(ekp[i][j]))
			for z := 0; z < len(ekp[i][j]); z++ {
				ekpb[i][j][z] = big.NewInt(ekp[i][j][z])
			}
		}
	}
	kc.EK = ekpb
	// Evaluator.
	eval := NewEvaluator(kc)
	// Ciphertext 0 (12345.678).
	c0p := [][]int64{{3928319131, 2322583730, 3128333473, -3588018546, -3023374777, -4583068212, 793865758, 2493685852, 4356875931, -1446151930, -155195535, 836686378, -2066582784, -4566179797, 2012557833, -4775884647, -2936070142, 2890040606, -3088406786, 1815033680, -2182161000, -2186917847, 35413812, 3547097307, 2678934739, -370357249, 3991359443, 2430971629, 454123781, -3417055595, 4371363907, 704824141},
		{-1556256741, -1792426624, 4406608444, -616705210, 2610504891, -1343850320, 1290341989, 3158043354, -3566746457, -3750192107, -4155437196, -1755722380, 1735554954, -2781807856, 2360617235, -3849070249, -3195929512, 2081665431, 178354185, 2509028358, 983465813, 1897902712, 1876620253, -2041701257, 4906960634, -642398707, 610816366, -3631039280, -2786954504, 636686208, 1293593149, -2004632067}}
	c0pb := make([][]*big.Int, len(c0p))
	for i := 0; i < len(c0p); i++ {
		c0pb[i] = make([]*big.Int, len(c0p[i]))
		for j := 0; j < len(c0p[i]); j++ {
			c0pb[i][j] = big.NewInt(c0p[i][j])
		}
	}
	// Ciphertext 1 (947.1273).
	c1p := [][]int64{{-801430505, -878485248, -3619214677, 3724339951, -4146097963, 2457285451, -166152169, 1146493397, -3930539301, -2148721601, 3525944698, 3336814609, 646191097, 1442584217, -1000306916, -1968228144, -3893801664, 4707109067, -4599778255, -1375893388, -3059390189, 7375620, 394482092, -1019907702, 1835725930, 3895685774, 4087024810, 1664162097, 4717028837, 4336772799, -4821294640, -1500042157},
		{4264193866, -1527646385, 1813200464, -3602499427, 1093465074, -258723476, 3334380716, 2120724897, 4787031003, 1767301708, -3677492885, 3839699240, -3074468330, -1120360217, 889645300, -4461608467, -2355083155, 3243450666, 1401312351, -4770263490, -1978232958, -4054237786, 1084734439, -422038801, -4464537862, -3270933427, -1790685853, 2585691139, 2152982520, 3343571564, 464216747, -4541369780}}
	c1pb := make([][]*big.Int, len(c1p))
	for i := 0; i < len(c1p); i++ {
		c1pb[i] = make([]*big.Int, len(c1p[i]))
		for j := 0; j < len(c1p[i]); j++ {
			c1pb[i][j] = big.NewInt(c1p[i][j])
		}
	}
	// Ciphertext multiplication.
	cr, err := eval.Mult(c0pb, c1pb)
	if err != nil {
		t.Errorf(err.Error())
	}
	// Check result.
	r := [][]int64{{787859912313, 1240006509703, -1822872053083, -793527086142, -684662907361, 2526846997639, 699411898497, 84390638574, 96430674797, -255402462867, -582522444013, 633294174711, 156110391411, -2756612685577, -864208953167, 1833959394090, -1630389832500, 1053863914378, -2014454942835, -428675979181, 1689153644999, 548362404868, -334516653875, -883426911269, 1763891634093, -293150022417, -1451092332874, 1852552201739, -658394181166, -813424555265, -630140763001, -866174559040},
		{1459368241335, 614961188533, 71369759170, -2008098050571, -685668442111, 346763477697, -33043828825, -614699121713, -4168501874970, 900149157498, -20251894743, -1774880426075, 197027302134, 482086032332, 305489583132, -323919229530, -513546606468, -713642274025, -3424835238928, 1810460169794, 676410066513, -1220581489852, -148875607383, -1030071930087, 504395304563, 781376694678, 1428434392519, 1287600027546, 531772649715, -1488003859943, 652784063365, -2344289619918}}
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[i]); j++ {
			if rb := big.NewInt(r[i][j]); rb.Cmp(cr[i][j]) != 0 {
				t.Errorf("expected %s for position [%d][%d] but got %s", rb.String(), i, j, cr[i][j].String())
				break
			}
		}
	}
}

// ######################################################################
// BFV FOR N = 1024, HERATION FOR N = 512
// ######################################################################
// TestBFVMult1024.

func TestBFVSAdd1024(t *testing.T) {
	// Create parameters.
	p, err := params.New(params.PLBFV1024)
	if err != nil {
		t.Error(err)
	}
	// SIM2D codec.
	sc, err := sim2d.New(p)
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := Setup("PLBFV1024.kc", o, p)
	if err != nil {
		t.Error(err)
	}
	// Cipher.
	cip, err := NewCipher(kc)
	if err != nil {
		t.Error(err)
	}
	// Evaluator.
	eval := NewEvaluator(kc)
	// Case: message (12345.678) + scalar (42.122).
	m := 12345.678
	s := 42.122
	// SIM2D encode.
	mpb := sc.Enc(m)
	spb := sc.Enc(s)
	// Encrypt message.
	c, err := cip.Enc(mpb)
	if err != nil {
		t.Error(err)
	}
	// Addition.
	cr, err := eval.SAdd(c, spb)
	if err != nil {
		t.Error(err)
	}
	// Decrypt.
	crd, err := cip.Dec(cr)
	if err != nil {
		t.Error(err)
	}
	// SIM2D decode.
	mrd, err := sc.Dec(crd)
	if err != nil {
		t.Error(err)
	}
	// Check result.
	if mr := m + s; mrd != mr {
		t.Errorf("expected %f for %f + %f, but got %f", mr, m, s, mrd)
	}
}

func TestBFVAdd1024(t *testing.T) {
	// Create parameters.
	p, err := params.New(params.PLBFV1024)
	if err != nil {
		t.Error(err)
	}
	// SIM2D codec.
	sc, err := sim2d.New(p)
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := Setup("PLBFV1024.kc", o, p)
	if err != nil {
		t.Error(err)
	}
	// Cipher.
	cip, err := NewCipher(kc)
	if err != nil {
		t.Error(err)
	}
	// Evaluator.
	eval := NewEvaluator(kc)
	// SIM2D encode.
	// Case: Message0 (12345.678) + Message1 (947.1273).
	m0 := params.M0
	m1 := params.M1
	m0pb := sc.Enc(m0)
	m1pb := sc.Enc(m1)
	// Encrypt messages.
	c0, err := cip.Enc(m0pb)
	if err != nil {
		t.Error(err)
	}
	c1, err := cip.Enc(m1pb)
	if err != nil {
		t.Error(err)
	}
	// Addition.
	cr := eval.Add(c0, c1)
	// Decrypt.
	crd, err := cip.Dec(cr)
	if err != nil {
		t.Error(err)
	}
	// SIM2D decode.
	mrd, err := sc.Dec(crd)
	if err != nil {
		t.Error(err)
	}
	// Check result.
	if mr := m0 + m1; mrd != mr {
		t.Errorf("expected %f for %f + %f, but got %f", mr, m0, m1, mrd)
	}
}

func TestBFVMult1024(t *testing.T) {
	// Create parameters.
	p, err := params.New(params.PLBFV1024)
	if err != nil {
		t.Error(err)
	}
	// SIM2D codec.
	sc, err := sim2d.New(p)
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := NewKeychain(o, p)
	if err != nil {
		t.Error(err)
	}
	// Cipher.
	cip, err := NewCipher(kc)
	if err != nil {
		t.Error(err)
	}
	// Evaluator.
	eval := NewEvaluator(kc)
	// Case: Messages.
	m0 := params.M2
	m1 := params.M3
	// SIM2D encode.
	m0pb := sc.Enc(m0)
	m1pb := sc.Enc(m1)
	// Encrypt messages.
	c0, err := cip.Enc(m0pb)
	if err != nil {
		t.Error(err)
	}
	c1, err := cip.Enc(m1pb)
	if err != nil {
		t.Error(err)
	}
	// Multiplication.
	cr, err := eval.Mult(c0, c1)
	if err != nil {
		t.Error(err)
	}
	// Decrypt.
	crd, err := cip.Dec(cr)
	if err != nil {
		t.Error(err)
	}
	// SIM2D decode.
	mrd, err := sc.Dec(crd)
	if err != nil {
		t.Error(err)
	}
	// Check result.
	if mr := m0 * m1; mrd != mr {
		t.Errorf("expected %f for %f x %f, but got %f", mr, m0, m1, mrd)
	}
}
