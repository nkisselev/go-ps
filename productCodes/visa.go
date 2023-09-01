package productCodes

// GetVisaProductName GetVisaProductCode returns the product name for the given product code
func GetVisaProductName(productCode string) string {
	var v, ex = visaProductCodes[productCode]
	if ex {
		return v
	}
	return ""
}

var visaProductCodes = map[string]string{
	"A":  "Visa Traditional",
	"AX": "American Express",
	"B":  "Visa Traditional Rewards",
	"C":  "Visa Signature",
	"D":  "Visa Signature Preferred",
	"DI": "Discover",
	"DN": "Diners",
	"E":  "Proprietary ATM",
	"F":  "Visa Classic",
	"G":  "Visa Business",
	"G1": "Visa Signature Business",
	"G2": "Visa Business Check Card",
	"G3": "Visa Business Enhanced",
	"G4": "Visa Infinite Business",
	"G5": "Visa Business Rewards",
	"I":  "Visa Infinite",
	"I1": "Visa Infinite Privilege",
	"I2": "Visa UHNW",
	"J3": "Visa Healthcare",
	"JC": "JCB",
	"K":  "Visa Corporate T&E",
	"K1": "Visa Government Corporate T&E",
	"L":  "Visa Electron",
	"M":  "MasterCard",
	"N":  "Visa Platinum",
	"N1": "Visa Rewards",
	"N2": "Visa Select",
	"P":  "Visa Gold",
	"Q":  "Private Label",
	"Q1": "Private Label Prepaid",
	"Q2": "Private Label Basic",
	"Q3": "Private Label Standard",
	"Q4": "Private Label Enhanced",
	"Q5": "Private Label Specialized",
	"Q6": "Private Label Premium",
	"R":  "Proprietary",
	"S":  "Visa Purchasing",
	"S1": "Visa Purchasing with Fleet",
	"S2": "Visa Government Purchasing",
	"S3": "Visa Government Purchasing with Fleet",
	"S4": "Visa Commercial Agriculture",
	"S5": "Visa Commercial Transport",
	"S6": "Visa Commercial Marketplace",
	"U":  "Visa Travel Money",
	"V":  "Visa V PAY",
}
