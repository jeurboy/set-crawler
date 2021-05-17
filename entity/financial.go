package entity

type FinancialPage struct {
	Title          string `pagser:"title"`
	FinancialTable struct {
		Title         DateHeaderRaw   `pagser:"thead tr th"`
		FinancialData []FinancialData `pagser:"tbody tr"`
	} `pagser:"table->eq(0)"`
}

//" วันที่\n    ราคาเปิด\n    ราคาสูงสุด\n    ราคาต่ำสุด\n    ราคาปิด\n    เปลี่ยนแปลง\n    %เปลี่ยนแปลง\n    ปริมาณรวม(หุ้น)\n    มูลค่ารวม('000 บาท)",
type FinancialData struct {
	Column []string `pagser:"td"`
}

// Assets (M.Baht)	2,746,714.79	2,544,182.88	2,484,438.68	2,355,483.87	2,232,314.16
// Liabilities (M.Baht)	1,343,556.82	1,258,337.79	1,183,399.06	1,036,989.52	983,758.42
// Equity (M.Baht)	953,735.25	882,040.40	878,604.11	875,083.83	818,671.75
// Paid-up Capital (M.Baht)	28,563.00	28,563.00	28,563.00	28,563.00	28,563.00
// Revenue (M.Baht)	492,191.01	1,633,977.16	2,239,718.62	2,353,090.01	2,023,990.58
// Net Profit/Net Loss (M.Baht)	32,587.61	37,765.81	92,950.60	119,683.94	135,179.60
// ROA (%)*	6.00	3.72	7.56	10.75	10.84
// ROE (%)*	7.90	4.29	10.60	14.13	17.09
// Net Profit Margin (%)	6.62	2.31	4.15	5.09	6.68

// สินทรัพย์ (ล้านบาท)	1,010.83	987.99	601.42	370.49	388.59
// หนี้สิน (ล้านบาท)	344.44	410.96	130.41	138.56	125.50
// ส่วนผู้ถือหุ้น (ล้านบาท)	530.38	471.95	471.01	231.93	263.09
// ทุนที่เรียกชำระแล้ว (ล้านบาท)	301.61	301.61	301.61	180.96	180.96
// รายได้ (ล้านบาท)	627.22	1,808.49	611.97	293.58	331.33
// กำไร/ขาดทุน (ล้านบาท)	58.43	-0.03	-146.29	-32.26	1.57
// ผลกำไรต่อสินทรัพย์ (%)*	10.79	-3.65	-27.49	-9.11	1.68
// ผลกำไรต่อส่วนผู้ถือหุ้น (%)*	11.66	-0.01	-41.62	-13.03	0.59
// อัตรากำไรสุทธิ (%)	9.32	-0.00	-23.90	-10.99	0.47
type FinancialRaw struct {
	Date            DateString
	Asset           DecimalString
	Liabilities     DecimalString
	Equity          DecimalString
	PaidUpCapital   DecimalString
	NetProfitOrLoss DecimalString
	ROA             DecimalString
	ROE             DecimalString
	NetProfitMargin DecimalString
}
