// Code generated from Turing.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 108, 783,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86,
	9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4, 91, 9,
	91, 4, 92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 4, 96, 9, 96,
	4, 97, 9, 97, 4, 98, 9, 98, 4, 99, 9, 99, 4, 100, 9, 100, 4, 101, 9, 101,
	4, 102, 9, 102, 4, 103, 9, 103, 4, 104, 9, 104, 4, 105, 9, 105, 4, 106,
	9, 106, 4, 107, 9, 107, 4, 108, 9, 108, 4, 109, 9, 109, 4, 110, 9, 110,
	4, 111, 9, 111, 4, 112, 9, 112, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 5, 5, 257, 10, 5, 3, 5, 3, 5, 3, 6, 6, 6, 262, 10, 6, 13, 6, 14, 6,
	263, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	5, 12, 298, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 5, 13, 311, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3,
	40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3,
	48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50,
	3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3,
	51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53,
	3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3,
	55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57,
	3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3,
	58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60,
	3, 60, 3, 61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 3,
	65, 3, 66, 3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 69,
	3, 70, 3, 70, 3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 73, 3, 73, 3, 74, 3,
	74, 3, 75, 3, 75, 3, 75, 3, 75, 3, 76, 3, 76, 3, 76, 3, 76, 3, 77, 3, 77,
	3, 77, 3, 77, 3, 78, 3, 78, 3, 78, 3, 79, 3, 79, 3, 80, 3, 80, 3, 81, 3,
	81, 3, 82, 3, 82, 3, 82, 3, 83, 3, 83, 3, 83, 3, 84, 3, 84, 3, 84, 3, 84,
	3, 84, 3, 85, 3, 85, 3, 85, 3, 85, 3, 86, 3, 86, 3, 86, 3, 87, 3, 87, 3,
	87, 3, 88, 3, 88, 3, 88, 3, 89, 3, 89, 3, 89, 3, 89, 3, 90, 3, 90, 3, 90,
	3, 90, 3, 91, 3, 91, 3, 91, 3, 91, 3, 92, 3, 92, 3, 92, 3, 93, 3, 93, 3,
	93, 3, 94, 3, 94, 3, 94, 3, 95, 3, 95, 3, 95, 3, 96, 3, 96, 3, 96, 3, 96,
	3, 96, 3, 97, 3, 97, 3, 97, 3, 97, 3, 97, 3, 98, 3, 98, 3, 98, 3, 98, 3,
	98, 3, 99, 3, 99, 3, 99, 3, 99, 3, 100, 3, 100, 3, 100, 3, 100, 3, 101,
	3, 101, 3, 101, 3, 101, 3, 101, 3, 102, 3, 102, 3, 102, 3, 102, 3, 102,
	3, 103, 3, 103, 3, 103, 3, 103, 3, 103, 3, 104, 3, 104, 3, 104, 7, 104,
	732, 10, 104, 12, 104, 14, 104, 735, 11, 104, 3, 105, 3, 105, 3, 106, 3,
	106, 3, 107, 6, 107, 742, 10, 107, 13, 107, 14, 107, 743, 3, 107, 3, 107,
	3, 108, 3, 108, 3, 108, 3, 108, 7, 108, 752, 10, 108, 12, 108, 14, 108,
	755, 11, 108, 3, 108, 3, 108, 3, 108, 3, 109, 3, 109, 7, 109, 762, 10,
	109, 12, 109, 14, 109, 765, 11, 109, 3, 110, 6, 110, 768, 10, 110, 13,
	110, 14, 110, 769, 3, 111, 3, 111, 3, 111, 3, 111, 3, 111, 3, 111, 3, 111,
	5, 111, 779, 10, 111, 3, 112, 3, 112, 3, 112, 3, 753, 2, 113, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43,
	23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61,
	32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79,
	41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97,
	50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56, 111, 57, 113,
	58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63, 125, 64, 127, 65, 129,
	66, 131, 67, 133, 68, 135, 69, 137, 70, 139, 71, 141, 72, 143, 73, 145,
	74, 147, 75, 149, 76, 151, 77, 153, 78, 155, 79, 157, 80, 159, 81, 161,
	82, 163, 83, 165, 84, 167, 85, 169, 86, 171, 87, 173, 88, 175, 89, 177,
	90, 179, 91, 181, 92, 183, 93, 185, 94, 187, 95, 189, 96, 191, 97, 193,
	98, 195, 99, 197, 100, 199, 101, 201, 102, 203, 103, 205, 104, 207, 105,
	209, 2, 211, 2, 213, 106, 215, 107, 217, 108, 219, 2, 221, 2, 223, 2, 3,
	2, 8, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15,
	34, 34, 4, 2, 12, 12, 15, 15, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 12,
	2, 36, 36, 41, 41, 65, 65, 94, 94, 99, 100, 104, 104, 112, 112, 116, 116,
	118, 118, 120, 120, 2, 790, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2,
	2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2,
	2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2,
	2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3,
	2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61,
	3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2,
	69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2,
	2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2,
	2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2,
	2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3,
	2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2,
	107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2,
	2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121,
	3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2,
	2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3,
	2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2, 141, 3, 2, 2, 2, 2,
	143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2, 2, 2, 2, 149, 3, 2,
	2, 2, 2, 151, 3, 2, 2, 2, 2, 153, 3, 2, 2, 2, 2, 155, 3, 2, 2, 2, 2, 157,
	3, 2, 2, 2, 2, 159, 3, 2, 2, 2, 2, 161, 3, 2, 2, 2, 2, 163, 3, 2, 2, 2,
	2, 165, 3, 2, 2, 2, 2, 167, 3, 2, 2, 2, 2, 169, 3, 2, 2, 2, 2, 171, 3,
	2, 2, 2, 2, 173, 3, 2, 2, 2, 2, 175, 3, 2, 2, 2, 2, 177, 3, 2, 2, 2, 2,
	179, 3, 2, 2, 2, 2, 181, 3, 2, 2, 2, 2, 183, 3, 2, 2, 2, 2, 185, 3, 2,
	2, 2, 2, 187, 3, 2, 2, 2, 2, 189, 3, 2, 2, 2, 2, 191, 3, 2, 2, 2, 2, 193,
	3, 2, 2, 2, 2, 195, 3, 2, 2, 2, 2, 197, 3, 2, 2, 2, 2, 199, 3, 2, 2, 2,
	2, 201, 3, 2, 2, 2, 2, 203, 3, 2, 2, 2, 2, 205, 3, 2, 2, 2, 2, 207, 3,
	2, 2, 2, 2, 213, 3, 2, 2, 2, 2, 215, 3, 2, 2, 2, 2, 217, 3, 2, 2, 2, 3,
	225, 3, 2, 2, 2, 5, 237, 3, 2, 2, 2, 7, 247, 3, 2, 2, 2, 9, 254, 3, 2,
	2, 2, 11, 261, 3, 2, 2, 2, 13, 265, 3, 2, 2, 2, 15, 269, 3, 2, 2, 2, 17,
	272, 3, 2, 2, 2, 19, 275, 3, 2, 2, 2, 21, 280, 3, 2, 2, 2, 23, 297, 3,
	2, 2, 2, 25, 310, 3, 2, 2, 2, 27, 312, 3, 2, 2, 2, 29, 318, 3, 2, 2, 2,
	31, 326, 3, 2, 2, 2, 33, 330, 3, 2, 2, 2, 35, 335, 3, 2, 2, 2, 37, 340,
	3, 2, 2, 2, 39, 343, 3, 2, 2, 2, 41, 348, 3, 2, 2, 2, 43, 354, 3, 2, 2,
	2, 45, 359, 3, 2, 2, 2, 47, 366, 3, 2, 2, 2, 49, 372, 3, 2, 2, 2, 51, 379,
	3, 2, 2, 2, 53, 386, 3, 2, 2, 2, 55, 390, 3, 2, 2, 2, 57, 395, 3, 2, 2,
	2, 59, 399, 3, 2, 2, 2, 61, 404, 3, 2, 2, 2, 63, 411, 3, 2, 2, 2, 65, 416,
	3, 2, 2, 2, 67, 422, 3, 2, 2, 2, 69, 427, 3, 2, 2, 2, 71, 437, 3, 2, 2,
	2, 73, 445, 3, 2, 2, 2, 75, 452, 3, 2, 2, 2, 77, 459, 3, 2, 2, 2, 79, 467,
	3, 2, 2, 2, 81, 477, 3, 2, 2, 2, 83, 480, 3, 2, 2, 2, 85, 484, 3, 2, 2,
	2, 87, 488, 3, 2, 2, 2, 89, 493, 3, 2, 2, 2, 91, 501, 3, 2, 2, 2, 93, 508,
	3, 2, 2, 2, 95, 514, 3, 2, 2, 2, 97, 518, 3, 2, 2, 2, 99, 525, 3, 2, 2,
	2, 101, 531, 3, 2, 2, 2, 103, 539, 3, 2, 2, 2, 105, 543, 3, 2, 2, 2, 107,
	548, 3, 2, 2, 2, 109, 553, 3, 2, 2, 2, 111, 559, 3, 2, 2, 2, 113, 564,
	3, 2, 2, 2, 115, 573, 3, 2, 2, 2, 117, 581, 3, 2, 2, 2, 119, 586, 3, 2,
	2, 2, 121, 590, 3, 2, 2, 2, 123, 592, 3, 2, 2, 2, 125, 594, 3, 2, 2, 2,
	127, 596, 3, 2, 2, 2, 129, 598, 3, 2, 2, 2, 131, 600, 3, 2, 2, 2, 133,
	603, 3, 2, 2, 2, 135, 605, 3, 2, 2, 2, 137, 607, 3, 2, 2, 2, 139, 610,
	3, 2, 2, 2, 141, 613, 3, 2, 2, 2, 143, 615, 3, 2, 2, 2, 145, 617, 3, 2,
	2, 2, 147, 619, 3, 2, 2, 2, 149, 621, 3, 2, 2, 2, 151, 625, 3, 2, 2, 2,
	153, 629, 3, 2, 2, 2, 155, 633, 3, 2, 2, 2, 157, 636, 3, 2, 2, 2, 159,
	638, 3, 2, 2, 2, 161, 640, 3, 2, 2, 2, 163, 642, 3, 2, 2, 2, 165, 645,
	3, 2, 2, 2, 167, 648, 3, 2, 2, 2, 169, 653, 3, 2, 2, 2, 171, 657, 3, 2,
	2, 2, 173, 660, 3, 2, 2, 2, 175, 663, 3, 2, 2, 2, 177, 666, 3, 2, 2, 2,
	179, 670, 3, 2, 2, 2, 181, 674, 3, 2, 2, 2, 183, 678, 3, 2, 2, 2, 185,
	681, 3, 2, 2, 2, 187, 684, 3, 2, 2, 2, 189, 687, 3, 2, 2, 2, 191, 690,
	3, 2, 2, 2, 193, 695, 3, 2, 2, 2, 195, 700, 3, 2, 2, 2, 197, 705, 3, 2,
	2, 2, 199, 709, 3, 2, 2, 2, 201, 713, 3, 2, 2, 2, 203, 718, 3, 2, 2, 2,
	205, 723, 3, 2, 2, 2, 207, 728, 3, 2, 2, 2, 209, 736, 3, 2, 2, 2, 211,
	738, 3, 2, 2, 2, 213, 741, 3, 2, 2, 2, 215, 747, 3, 2, 2, 2, 217, 759,
	3, 2, 2, 2, 219, 767, 3, 2, 2, 2, 221, 778, 3, 2, 2, 2, 223, 780, 3, 2,
	2, 2, 225, 226, 7, 119, 2, 2, 226, 227, 7, 112, 2, 2, 227, 228, 7, 115,
	2, 2, 228, 229, 7, 119, 2, 2, 229, 230, 7, 99, 2, 2, 230, 231, 7, 110,
	2, 2, 231, 232, 7, 107, 2, 2, 232, 233, 7, 104, 2, 2, 233, 234, 7, 107,
	2, 2, 234, 235, 7, 103, 2, 2, 235, 236, 7, 102, 2, 2, 236, 4, 3, 2, 2,
	2, 237, 238, 7, 114, 2, 2, 238, 239, 7, 103, 2, 2, 239, 240, 7, 116, 2,
	2, 240, 241, 7, 120, 2, 2, 241, 242, 7, 99, 2, 2, 242, 243, 7, 117, 2,
	2, 243, 244, 7, 107, 2, 2, 244, 245, 7, 120, 2, 2, 245, 246, 7, 103, 2,
	2, 246, 6, 3, 2, 2, 2, 247, 248, 7, 113, 2, 2, 248, 249, 7, 114, 2, 2,
	249, 250, 7, 99, 2, 2, 250, 251, 7, 115, 2, 2, 251, 252, 7, 119, 2, 2,
	252, 253, 7, 103, 2, 2, 253, 8, 3, 2, 2, 2, 254, 256, 7, 36, 2, 2, 255,
	257, 5, 219, 110, 2, 256, 255, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 258,
	3, 2, 2, 2, 258, 259, 7, 36, 2, 2, 259, 10, 3, 2, 2, 2, 260, 262, 5, 211,
	106, 2, 261, 260, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 261, 3, 2, 2,
	2, 263, 264, 3, 2, 2, 2, 264, 12, 3, 2, 2, 2, 265, 266, 7, 103, 2, 2, 266,
	267, 7, 112, 2, 2, 267, 268, 7, 102, 2, 2, 268, 14, 3, 2, 2, 2, 269, 270,
	7, 113, 2, 2, 270, 271, 7, 104, 2, 2, 271, 16, 3, 2, 2, 2, 272, 273, 7,
	118, 2, 2, 273, 274, 7, 113, 2, 2, 274, 18, 3, 2, 2, 2, 275, 276, 7, 118,
	2, 2, 276, 277, 7, 123, 2, 2, 277, 278, 7, 114, 2, 2, 278, 279, 7, 103,
	2, 2, 279, 20, 3, 2, 2, 2, 280, 281, 7, 120, 2, 2, 281, 282, 7, 99, 2,
	2, 282, 283, 7, 116, 2, 2, 283, 22, 3, 2, 2, 2, 284, 285, 7, 114, 2, 2,
	285, 286, 7, 116, 2, 2, 286, 287, 7, 113, 2, 2, 287, 288, 7, 101, 2, 2,
	288, 289, 7, 103, 2, 2, 289, 290, 7, 102, 2, 2, 290, 291, 7, 119, 2, 2,
	291, 292, 7, 116, 2, 2, 292, 298, 7, 103, 2, 2, 293, 294, 7, 114, 2, 2,
	294, 295, 7, 116, 2, 2, 295, 296, 7, 113, 2, 2, 296, 298, 7, 101, 2, 2,
	297, 284, 3, 2, 2, 2, 297, 293, 3, 2, 2, 2, 298, 24, 3, 2, 2, 2, 299, 300,
	7, 104, 2, 2, 300, 301, 7, 119, 2, 2, 301, 302, 7, 112, 2, 2, 302, 303,
	7, 101, 2, 2, 303, 304, 7, 118, 2, 2, 304, 305, 7, 107, 2, 2, 305, 306,
	7, 113, 2, 2, 306, 311, 7, 112, 2, 2, 307, 308, 7, 104, 2, 2, 308, 309,
	7, 101, 2, 2, 309, 311, 7, 112, 2, 2, 310, 299, 3, 2, 2, 2, 310, 307, 3,
	2, 2, 2, 311, 26, 3, 2, 2, 2, 312, 313, 7, 101, 2, 2, 313, 314, 7, 110,
	2, 2, 314, 315, 7, 99, 2, 2, 315, 316, 7, 117, 2, 2, 316, 317, 7, 117,
	2, 2, 317, 28, 3, 2, 2, 2, 318, 319, 7, 114, 2, 2, 319, 320, 7, 116, 2,
	2, 320, 321, 7, 113, 2, 2, 321, 322, 7, 101, 2, 2, 322, 323, 7, 103, 2,
	2, 323, 324, 7, 117, 2, 2, 324, 325, 7, 117, 2, 2, 325, 30, 3, 2, 2, 2,
	326, 327, 7, 104, 2, 2, 327, 328, 7, 113, 2, 2, 328, 329, 7, 116, 2, 2,
	329, 32, 3, 2, 2, 2, 330, 331, 7, 110, 2, 2, 331, 332, 7, 113, 2, 2, 332,
	333, 7, 113, 2, 2, 333, 334, 7, 114, 2, 2, 334, 34, 3, 2, 2, 2, 335, 336,
	7, 103, 2, 2, 336, 337, 7, 122, 2, 2, 337, 338, 7, 107, 2, 2, 338, 339,
	7, 118, 2, 2, 339, 36, 3, 2, 2, 2, 340, 341, 7, 107, 2, 2, 341, 342, 7,
	104, 2, 2, 342, 38, 3, 2, 2, 2, 343, 344, 7, 103, 2, 2, 344, 345, 7, 110,
	2, 2, 345, 346, 7, 117, 2, 2, 346, 347, 7, 103, 2, 2, 347, 40, 3, 2, 2,
	2, 348, 349, 7, 103, 2, 2, 349, 350, 7, 110, 2, 2, 350, 351, 7, 117, 2,
	2, 351, 352, 7, 107, 2, 2, 352, 353, 7, 104, 2, 2, 353, 42, 3, 2, 2, 2,
	354, 355, 7, 101, 2, 2, 355, 356, 7, 99, 2, 2, 356, 357, 7, 117, 2, 2,
	357, 358, 7, 103, 2, 2, 358, 44, 3, 2, 2, 2, 359, 360, 7, 99, 2, 2, 360,
	361, 7, 117, 2, 2, 361, 362, 7, 117, 2, 2, 362, 363, 7, 103, 2, 2, 363,
	364, 7, 116, 2, 2, 364, 365, 7, 118, 2, 2, 365, 46, 3, 2, 2, 2, 366, 367,
	7, 100, 2, 2, 367, 368, 7, 103, 2, 2, 368, 369, 7, 105, 2, 2, 369, 370,
	7, 107, 2, 2, 370, 371, 7, 112, 2, 2, 371, 48, 3, 2, 2, 2, 372, 373, 7,
	116, 2, 2, 373, 374, 7, 103, 2, 2, 374, 375, 7, 118, 2, 2, 375, 376, 7,
	119, 2, 2, 376, 377, 7, 116, 2, 2, 377, 378, 7, 112, 2, 2, 378, 50, 3,
	2, 2, 2, 379, 380, 7, 116, 2, 2, 380, 381, 7, 103, 2, 2, 381, 382, 7, 117,
	2, 2, 382, 383, 7, 119, 2, 2, 383, 384, 7, 110, 2, 2, 384, 385, 7, 118,
	2, 2, 385, 52, 3, 2, 2, 2, 386, 387, 7, 112, 2, 2, 387, 388, 7, 103, 2,
	2, 388, 389, 7, 121, 2, 2, 389, 54, 3, 2, 2, 2, 390, 391, 7, 104, 2, 2,
	391, 392, 7, 116, 2, 2, 392, 393, 7, 103, 2, 2, 393, 394, 7, 103, 2, 2,
	394, 56, 3, 2, 2, 2, 395, 396, 7, 118, 2, 2, 396, 397, 7, 99, 2, 2, 397,
	398, 7, 105, 2, 2, 398, 58, 3, 2, 2, 2, 399, 400, 7, 104, 2, 2, 400, 401,
	7, 113, 2, 2, 401, 402, 7, 116, 2, 2, 402, 403, 7, 109, 2, 2, 403, 60,
	3, 2, 2, 2, 404, 405, 7, 117, 2, 2, 405, 406, 7, 107, 2, 2, 406, 407, 7,
	105, 2, 2, 407, 408, 7, 112, 2, 2, 408, 409, 7, 99, 2, 2, 409, 410, 7,
	110, 2, 2, 410, 62, 3, 2, 2, 2, 411, 412, 7, 121, 2, 2, 412, 413, 7, 99,
	2, 2, 413, 414, 7, 107, 2, 2, 414, 415, 7, 118, 2, 2, 415, 64, 3, 2, 2,
	2, 416, 417, 7, 114, 2, 2, 417, 418, 7, 99, 2, 2, 418, 419, 7, 119, 2,
	2, 419, 420, 7, 117, 2, 2, 420, 421, 7, 103, 2, 2, 421, 66, 3, 2, 2, 2,
	422, 423, 7, 115, 2, 2, 423, 424, 7, 119, 2, 2, 424, 425, 7, 107, 2, 2,
	425, 426, 7, 118, 2, 2, 426, 68, 3, 2, 2, 2, 427, 428, 7, 119, 2, 2, 428,
	429, 7, 112, 2, 2, 429, 430, 7, 101, 2, 2, 430, 431, 7, 106, 2, 2, 431,
	432, 7, 103, 2, 2, 432, 433, 7, 101, 2, 2, 433, 434, 7, 109, 2, 2, 434,
	435, 7, 103, 2, 2, 435, 436, 7, 102, 2, 2, 436, 70, 3, 2, 2, 2, 437, 438,
	7, 101, 2, 2, 438, 439, 7, 106, 2, 2, 439, 440, 7, 103, 2, 2, 440, 441,
	7, 101, 2, 2, 441, 442, 7, 109, 2, 2, 442, 443, 7, 103, 2, 2, 443, 444,
	7, 102, 2, 2, 444, 72, 3, 2, 2, 2, 445, 446, 7, 103, 2, 2, 446, 447, 7,
	122, 2, 2, 447, 448, 7, 114, 2, 2, 448, 449, 7, 113, 2, 2, 449, 450, 7,
	116, 2, 2, 450, 451, 7, 118, 2, 2, 451, 74, 3, 2, 2, 2, 452, 453, 7, 107,
	2, 2, 453, 454, 7, 111, 2, 2, 454, 455, 7, 114, 2, 2, 455, 456, 7, 113,
	2, 2, 456, 457, 7, 116, 2, 2, 457, 458, 7, 118, 2, 2, 458, 76, 3, 2, 2,
	2, 459, 460, 7, 107, 2, 2, 460, 461, 7, 112, 2, 2, 461, 462, 7, 106, 2,
	2, 462, 463, 7, 103, 2, 2, 463, 464, 7, 116, 2, 2, 464, 465, 7, 107, 2,
	2, 465, 466, 7, 118, 2, 2, 466, 78, 3, 2, 2, 2, 467, 468, 7, 107, 2, 2,
	468, 469, 7, 111, 2, 2, 469, 470, 7, 114, 2, 2, 470, 471, 7, 110, 2, 2,
	471, 472, 7, 103, 2, 2, 472, 473, 7, 111, 2, 2, 473, 474, 7, 103, 2, 2,
	474, 475, 7, 112, 2, 2, 475, 476, 7, 118, 2, 2, 476, 80, 3, 2, 2, 2, 477,
	478, 7, 100, 2, 2, 478, 479, 7, 123, 2, 2, 479, 82, 3, 2, 2, 2, 480, 481,
	7, 114, 2, 2, 481, 482, 7, 119, 2, 2, 482, 483, 7, 118, 2, 2, 483, 84,
	3, 2, 2, 2, 484, 485, 7, 107, 2, 2, 485, 486, 7, 112, 2, 2, 486, 487, 7,
	118, 2, 2, 487, 86, 3, 2, 2, 2, 488, 489, 7, 116, 2, 2, 489, 490, 7, 103,
	2, 2, 490, 491, 7, 99, 2, 2, 491, 492, 7, 110, 2, 2, 492, 88, 3, 2, 2,
	2, 493, 494, 7, 100, 2, 2, 494, 495, 7, 113, 2, 2, 495, 496, 7, 113, 2,
	2, 496, 497, 7, 110, 2, 2, 497, 498, 7, 103, 2, 2, 498, 499, 7, 99, 2,
	2, 499, 500, 7, 112, 2, 2, 500, 90, 3, 2, 2, 2, 501, 502, 7, 117, 2, 2,
	502, 503, 7, 118, 2, 2, 503, 504, 7, 116, 2, 2, 504, 505, 7, 107, 2, 2,
	505, 506, 7, 112, 2, 2, 506, 507, 7, 105, 2, 2, 507, 92, 3, 2, 2, 2, 508,
	509, 7, 99, 2, 2, 509, 510, 7, 116, 2, 2, 510, 511, 7, 116, 2, 2, 511,
	512, 7, 99, 2, 2, 512, 513, 7, 123, 2, 2, 513, 94, 3, 2, 2, 2, 514, 515,
	7, 117, 2, 2, 515, 516, 7, 103, 2, 2, 516, 517, 7, 118, 2, 2, 517, 96,
	3, 2, 2, 2, 518, 519, 7, 116, 2, 2, 519, 520, 7, 103, 2, 2, 520, 521, 7,
	101, 2, 2, 521, 522, 7, 113, 2, 2, 522, 523, 7, 116, 2, 2, 523, 524, 7,
	102, 2, 2, 524, 98, 3, 2, 2, 2, 525, 526, 7, 119, 2, 2, 526, 527, 7, 112,
	2, 2, 527, 528, 7, 107, 2, 2, 528, 529, 7, 113, 2, 2, 529, 530, 7, 112,
	2, 2, 530, 100, 3, 2, 2, 2, 531, 532, 7, 114, 2, 2, 532, 533, 7, 113, 2,
	2, 533, 534, 7, 107, 2, 2, 534, 535, 7, 112, 2, 2, 535, 536, 7, 118, 2,
	2, 536, 537, 7, 103, 2, 2, 537, 538, 7, 116, 2, 2, 538, 102, 3, 2, 2, 2,
	539, 540, 7, 112, 2, 2, 540, 541, 7, 99, 2, 2, 541, 542, 7, 118, 2, 2,
	542, 104, 3, 2, 2, 2, 543, 544, 7, 107, 2, 2, 544, 545, 7, 112, 2, 2, 545,
	546, 7, 118, 2, 2, 546, 547, 7, 112, 2, 2, 547, 106, 3, 2, 2, 2, 548, 549,
	7, 112, 2, 2, 549, 550, 7, 99, 2, 2, 550, 551, 7, 118, 2, 2, 551, 552,
	7, 112, 2, 2, 552, 108, 3, 2, 2, 2, 553, 554, 7, 116, 2, 2, 554, 555, 7,
	103, 2, 2, 555, 556, 7, 99, 2, 2, 556, 557, 7, 110, 2, 2, 557, 558, 7,
	112, 2, 2, 558, 110, 3, 2, 2, 2, 559, 560, 7, 101, 2, 2, 560, 561, 7, 106,
	2, 2, 561, 562, 7, 99, 2, 2, 562, 563, 7, 116, 2, 2, 563, 112, 3, 2, 2,
	2, 564, 565, 7, 102, 2, 2, 565, 566, 7, 103, 2, 2, 566, 567, 7, 104, 2,
	2, 567, 568, 7, 103, 2, 2, 568, 569, 7, 116, 2, 2, 569, 570, 7, 116, 2,
	2, 570, 571, 7, 103, 2, 2, 571, 572, 7, 102, 2, 2, 572, 114, 3, 2, 2, 2,
	573, 574, 7, 104, 2, 2, 574, 575, 7, 113, 2, 2, 575, 576, 7, 116, 2, 2,
	576, 577, 7, 121, 2, 2, 577, 578, 7, 99, 2, 2, 578, 579, 7, 116, 2, 2,
	579, 580, 7, 102, 2, 2, 580, 116, 3, 2, 2, 2, 581, 582, 7, 100, 2, 2, 582,
	583, 7, 113, 2, 2, 583, 584, 7, 102, 2, 2, 584, 585, 7, 123, 2, 2, 585,
	118, 3, 2, 2, 2, 586, 587, 7, 112, 2, 2, 587, 588, 7, 113, 2, 2, 588, 589,
	7, 118, 2, 2, 589, 120, 3, 2, 2, 2, 590, 591, 7, 96, 2, 2, 591, 122, 3,
	2, 2, 2, 592, 593, 7, 60, 2, 2, 593, 124, 3, 2, 2, 2, 594, 595, 7, 42,
	2, 2, 595, 126, 3, 2, 2, 2, 596, 597, 7, 43, 2, 2, 597, 128, 3, 2, 2, 2,
	598, 599, 7, 48, 2, 2, 599, 130, 3, 2, 2, 2, 600, 601, 7, 48, 2, 2, 601,
	602, 7, 48, 2, 2, 602, 132, 3, 2, 2, 2, 603, 604, 7, 46, 2, 2, 604, 134,
	3, 2, 2, 2, 605, 606, 7, 37, 2, 2, 606, 136, 3, 2, 2, 2, 607, 608, 7, 47,
	2, 2, 608, 609, 7, 64, 2, 2, 609, 138, 3, 2, 2, 2, 610, 611, 7, 60, 2,
	2, 611, 612, 7, 63, 2, 2, 612, 140, 3, 2, 2, 2, 613, 614, 7, 45, 2, 2,
	614, 142, 3, 2, 2, 2, 615, 616, 7, 47, 2, 2, 616, 144, 3, 2, 2, 2, 617,
	618, 7, 44, 2, 2, 618, 146, 3, 2, 2, 2, 619, 620, 7, 49, 2, 2, 620, 148,
	3, 2, 2, 2, 621, 622, 7, 102, 2, 2, 622, 623, 7, 107, 2, 2, 623, 624, 7,
	120, 2, 2, 624, 150, 3, 2, 2, 2, 625, 626, 7, 111, 2, 2, 626, 627, 7, 113,
	2, 2, 627, 628, 7, 102, 2, 2, 628, 152, 3, 2, 2, 2, 629, 630, 7, 116, 2,
	2, 630, 631, 7, 103, 2, 2, 631, 632, 7, 111, 2, 2, 632, 154, 3, 2, 2, 2,
	633, 634, 7, 44, 2, 2, 634, 635, 7, 44, 2, 2, 635, 156, 3, 2, 2, 2, 636,
	637, 7, 62, 2, 2, 637, 158, 3, 2, 2, 2, 638, 639, 7, 64, 2, 2, 639, 160,
	3, 2, 2, 2, 640, 641, 7, 63, 2, 2, 641, 162, 3, 2, 2, 2, 642, 643, 7, 62,
	2, 2, 643, 644, 7, 63, 2, 2, 644, 164, 3, 2, 2, 2, 645, 646, 7, 64, 2,
	2, 646, 647, 7, 63, 2, 2, 647, 166, 3, 2, 2, 2, 648, 649, 7, 112, 2, 2,
	649, 650, 7, 113, 2, 2, 650, 651, 7, 118, 2, 2, 651, 652, 7, 63, 2, 2,
	652, 168, 3, 2, 2, 2, 653, 654, 7, 99, 2, 2, 654, 655, 7, 112, 2, 2, 655,
	656, 7, 102, 2, 2, 656, 170, 3, 2, 2, 2, 657, 658, 7, 113, 2, 2, 658, 659,
	7, 116, 2, 2, 659, 172, 3, 2, 2, 2, 660, 661, 7, 63, 2, 2, 661, 662, 7,
	64, 2, 2, 662, 174, 3, 2, 2, 2, 663, 664, 7, 107, 2, 2, 664, 665, 7, 112,
	2, 2, 665, 176, 3, 2, 2, 2, 666, 667, 7, 117, 2, 2, 667, 668, 7, 106, 2,
	2, 668, 669, 7, 116, 2, 2, 669, 178, 3, 2, 2, 2, 670, 671, 7, 117, 2, 2,
	671, 672, 7, 106, 2, 2, 672, 673, 7, 110, 2, 2, 673, 180, 3, 2, 2, 2, 674,
	675, 7, 122, 2, 2, 675, 676, 7, 113, 2, 2, 676, 677, 7, 116, 2, 2, 677,
	182, 3, 2, 2, 2, 678, 679, 7, 45, 2, 2, 679, 680, 7, 63, 2, 2, 680, 184,
	3, 2, 2, 2, 681, 682, 7, 47, 2, 2, 682, 683, 7, 63, 2, 2, 683, 186, 3,
	2, 2, 2, 684, 685, 7, 44, 2, 2, 685, 686, 7, 63, 2, 2, 686, 188, 3, 2,
	2, 2, 687, 688, 7, 49, 2, 2, 688, 689, 7, 63, 2, 2, 689, 190, 3, 2, 2,
	2, 690, 691, 7, 102, 2, 2, 691, 692, 7, 107, 2, 2, 692, 693, 7, 120, 2,
	2, 693, 694, 7, 63, 2, 2, 694, 192, 3, 2, 2, 2, 695, 696, 7, 111, 2, 2,
	696, 697, 7, 113, 2, 2, 697, 698, 7, 102, 2, 2, 698, 699, 7, 63, 2, 2,
	699, 194, 3, 2, 2, 2, 700, 701, 7, 99, 2, 2, 701, 702, 7, 112, 2, 2, 702,
	703, 7, 102, 2, 2, 703, 704, 7, 63, 2, 2, 704, 196, 3, 2, 2, 2, 705, 706,
	7, 113, 2, 2, 706, 707, 7, 116, 2, 2, 707, 708, 7, 63, 2, 2, 708, 198,
	3, 2, 2, 2, 709, 710, 7, 63, 2, 2, 710, 711, 7, 64, 2, 2, 711, 712, 7,
	63, 2, 2, 712, 200, 3, 2, 2, 2, 713, 714, 7, 117, 2, 2, 714, 715, 7, 106,
	2, 2, 715, 716, 7, 116, 2, 2, 716, 717, 7, 63, 2, 2, 717, 202, 3, 2, 2,
	2, 718, 719, 7, 117, 2, 2, 719, 720, 7, 106, 2, 2, 720, 721, 7, 110, 2,
	2, 721, 722, 7, 63, 2, 2, 722, 204, 3, 2, 2, 2, 723, 724, 7, 122, 2, 2,
	724, 725, 7, 113, 2, 2, 725, 726, 7, 116, 2, 2, 726, 727, 7, 63, 2, 2,
	727, 206, 3, 2, 2, 2, 728, 733, 5, 209, 105, 2, 729, 732, 5, 209, 105,
	2, 730, 732, 5, 211, 106, 2, 731, 729, 3, 2, 2, 2, 731, 730, 3, 2, 2, 2,
	732, 735, 3, 2, 2, 2, 733, 731, 3, 2, 2, 2, 733, 734, 3, 2, 2, 2, 734,
	208, 3, 2, 2, 2, 735, 733, 3, 2, 2, 2, 736, 737, 9, 2, 2, 2, 737, 210,
	3, 2, 2, 2, 738, 739, 9, 3, 2, 2, 739, 212, 3, 2, 2, 2, 740, 742, 9, 4,
	2, 2, 741, 740, 3, 2, 2, 2, 742, 743, 3, 2, 2, 2, 743, 741, 3, 2, 2, 2,
	743, 744, 3, 2, 2, 2, 744, 745, 3, 2, 2, 2, 745, 746, 8, 107, 2, 2, 746,
	214, 3, 2, 2, 2, 747, 748, 7, 49, 2, 2, 748, 749, 7, 44, 2, 2, 749, 753,
	3, 2, 2, 2, 750, 752, 11, 2, 2, 2, 751, 750, 3, 2, 2, 2, 752, 755, 3, 2,
	2, 2, 753, 754, 3, 2, 2, 2, 753, 751, 3, 2, 2, 2, 754, 756, 3, 2, 2, 2,
	755, 753, 3, 2, 2, 2, 756, 757, 7, 44, 2, 2, 757, 758, 7, 49, 2, 2, 758,
	216, 3, 2, 2, 2, 759, 763, 7, 39, 2, 2, 760, 762, 10, 5, 2, 2, 761, 760,
	3, 2, 2, 2, 762, 765, 3, 2, 2, 2, 763, 761, 3, 2, 2, 2, 763, 764, 3, 2,
	2, 2, 764, 218, 3, 2, 2, 2, 765, 763, 3, 2, 2, 2, 766, 768, 5, 221, 111,
	2, 767, 766, 3, 2, 2, 2, 768, 769, 3, 2, 2, 2, 769, 767, 3, 2, 2, 2, 769,
	770, 3, 2, 2, 2, 770, 220, 3, 2, 2, 2, 771, 779, 10, 6, 2, 2, 772, 779,
	5, 223, 112, 2, 773, 774, 7, 94, 2, 2, 774, 779, 7, 12, 2, 2, 775, 776,
	7, 94, 2, 2, 776, 777, 7, 15, 2, 2, 777, 779, 7, 12, 2, 2, 778, 771, 3,
	2, 2, 2, 778, 772, 3, 2, 2, 2, 778, 773, 3, 2, 2, 2, 778, 775, 3, 2, 2,
	2, 779, 222, 3, 2, 2, 2, 780, 781, 7, 94, 2, 2, 781, 782, 9, 7, 2, 2, 782,
	224, 3, 2, 2, 2, 14, 2, 256, 263, 297, 310, 731, 733, 743, 753, 763, 769,
	778, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'unqualified'", "'pervasive'", "'opaque'", "", "", "'end'", "'of'",
	"'to'", "'type'", "'var'", "", "", "'class'", "'process'", "'for'", "'loop'",
	"'exit'", "'if'", "'else'", "'elsif'", "'case'", "'assert'", "'begin'",
	"'return'", "'result'", "'new'", "'free'", "'tag'", "'fork'", "'signal'",
	"'wait'", "'pause'", "'quit'", "'unchecked'", "'checked'", "'export'",
	"'import'", "'inherit'", "'implement'", "'by'", "'put'", "'int'", "'real'",
	"'boolean'", "'string'", "'array'", "'set'", "'record'", "'union'", "'pointer'",
	"'nat'", "'intn'", "'natn'", "'realn'", "'char'", "'deferred'", "'forward'",
	"'body'", "'not'", "'^'", "':'", "'('", "')'", "'.'", "'..'", "','", "'#'",
	"'->'", "':='", "'+'", "'-'", "'*'", "'/'", "'div'", "'mod'", "'rem'",
	"'**'", "'<'", "'>'", "'='", "'<='", "'>='", "'not='", "'and'", "'or'",
	"'=>'", "'in'", "'shr'", "'shl'", "'xor'", "'+='", "'-='", "'*='", "'/='",
	"'div='", "'mod='", "'and='", "'or='", "'=>='", "'shr='", "'shl='", "'xor='",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "STRING_LITERAL", "INTEGER_LITERAL", "END", "OF", "TO",
	"TYPE", "VAR", "PROCEDURE", "FUNCTION", "CLASS", "PROCESS", "FOR", "LOOP",
	"EXIT", "IF", "ELSE", "ELSIF", "CASE", "ASSERT", "BEGIN", "RETURN", "RESULT",
	"NEW", "FREE", "TAG", "FORK", "SIGNAL", "WAIT", "PAUSE", "QUIT", "UNCHECKED",
	"CHECKED", "EXPORT", "IMPORT", "INHERIT", "IMPLEMENT", "BY", "PUT", "INT",
	"REAL", "BOOLEAN", "STRING", "ARRAY", "SET", "RECORD", "UNION", "POINTER",
	"NAT", "INTN", "NATN", "REALN", "CHAR", "DEFERRED", "FORWARD", "BODY",
	"NOT", "CARET", "COLON", "L_PAREN", "R_PAREN", "DOT", "RANGE", "COMMA",
	"CHEAT", "DEREFERENCE", "ASSIGNMENT", "PLUS", "MINUS", "MULTIPLY", "DIVIDE",
	"DIV", "MOD", "REM", "EXP", "LESSTHAN", "GREATERTHAN", "EQUAL", "LESSTHANEQUAL",
	"GREATERTHANEQUAL", "NOTEQUAL", "AND", "OR", "IMPLIES", "IN", "SHR", "SHL",
	"XOR", "PLUSEQUALS", "MINUSEQUALS", "MULTIPLYEQUALS", "DIVIDEEQUALS", "DIVEQUALS",
	"MODEQUALS", "ANDEQUALS", "OREQUALS", "IMPLIESEQUALS", "SHREQUALS", "SHLEQUALS",
	"XOREQUALS", "IDENTIFIER", "WHITESPACE", "BLOCK_COMMENT", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "STRING_LITERAL", "INTEGER_LITERAL", "END", "OF",
	"TO", "TYPE", "VAR", "PROCEDURE", "FUNCTION", "CLASS", "PROCESS", "FOR",
	"LOOP", "EXIT", "IF", "ELSE", "ELSIF", "CASE", "ASSERT", "BEGIN", "RETURN",
	"RESULT", "NEW", "FREE", "TAG", "FORK", "SIGNAL", "WAIT", "PAUSE", "QUIT",
	"UNCHECKED", "CHECKED", "EXPORT", "IMPORT", "INHERIT", "IMPLEMENT", "BY",
	"PUT", "INT", "REAL", "BOOLEAN", "STRING", "ARRAY", "SET", "RECORD", "UNION",
	"POINTER", "NAT", "INTN", "NATN", "REALN", "CHAR", "DEFERRED", "FORWARD",
	"BODY", "NOT", "CARET", "COLON", "L_PAREN", "R_PAREN", "DOT", "RANGE",
	"COMMA", "CHEAT", "DEREFERENCE", "ASSIGNMENT", "PLUS", "MINUS", "MULTIPLY",
	"DIVIDE", "DIV", "MOD", "REM", "EXP", "LESSTHAN", "GREATERTHAN", "EQUAL",
	"LESSTHANEQUAL", "GREATERTHANEQUAL", "NOTEQUAL", "AND", "OR", "IMPLIES",
	"IN", "SHR", "SHL", "XOR", "PLUSEQUALS", "MINUSEQUALS", "MULTIPLYEQUALS",
	"DIVIDEEQUALS", "DIVEQUALS", "MODEQUALS", "ANDEQUALS", "OREQUALS", "IMPLIESEQUALS",
	"SHREQUALS", "SHLEQUALS", "XOREQUALS", "IDENTIFIER", "NON_DIGIT", "DIGIT",
	"WHITESPACE", "BLOCK_COMMENT", "LINE_COMMENT", "STRING_CHAR_SEQUENCE",
	"STRING_CHAR", "ESCAPE_SEQUENCE",
}

type TuringLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewTuringLexer(input antlr.CharStream) *TuringLexer {

	l := new(TuringLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Turing.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TuringLexer tokens.
const (
	TuringLexerT__0             = 1
	TuringLexerT__1             = 2
	TuringLexerT__2             = 3
	TuringLexerSTRING_LITERAL   = 4
	TuringLexerINTEGER_LITERAL  = 5
	TuringLexerEND              = 6
	TuringLexerOF               = 7
	TuringLexerTO               = 8
	TuringLexerTYPE             = 9
	TuringLexerVAR              = 10
	TuringLexerPROCEDURE        = 11
	TuringLexerFUNCTION         = 12
	TuringLexerCLASS            = 13
	TuringLexerPROCESS          = 14
	TuringLexerFOR              = 15
	TuringLexerLOOP             = 16
	TuringLexerEXIT             = 17
	TuringLexerIF               = 18
	TuringLexerELSE             = 19
	TuringLexerELSIF            = 20
	TuringLexerCASE             = 21
	TuringLexerASSERT           = 22
	TuringLexerBEGIN            = 23
	TuringLexerRETURN           = 24
	TuringLexerRESULT           = 25
	TuringLexerNEW              = 26
	TuringLexerFREE             = 27
	TuringLexerTAG              = 28
	TuringLexerFORK             = 29
	TuringLexerSIGNAL           = 30
	TuringLexerWAIT             = 31
	TuringLexerPAUSE            = 32
	TuringLexerQUIT             = 33
	TuringLexerUNCHECKED        = 34
	TuringLexerCHECKED          = 35
	TuringLexerEXPORT           = 36
	TuringLexerIMPORT           = 37
	TuringLexerINHERIT          = 38
	TuringLexerIMPLEMENT        = 39
	TuringLexerBY               = 40
	TuringLexerPUT              = 41
	TuringLexerINT              = 42
	TuringLexerREAL             = 43
	TuringLexerBOOLEAN          = 44
	TuringLexerSTRING           = 45
	TuringLexerARRAY            = 46
	TuringLexerSET              = 47
	TuringLexerRECORD           = 48
	TuringLexerUNION            = 49
	TuringLexerPOINTER          = 50
	TuringLexerNAT              = 51
	TuringLexerINTN             = 52
	TuringLexerNATN             = 53
	TuringLexerREALN            = 54
	TuringLexerCHAR             = 55
	TuringLexerDEFERRED         = 56
	TuringLexerFORWARD          = 57
	TuringLexerBODY             = 58
	TuringLexerNOT              = 59
	TuringLexerCARET            = 60
	TuringLexerCOLON            = 61
	TuringLexerL_PAREN          = 62
	TuringLexerR_PAREN          = 63
	TuringLexerDOT              = 64
	TuringLexerRANGE            = 65
	TuringLexerCOMMA            = 66
	TuringLexerCHEAT            = 67
	TuringLexerDEREFERENCE      = 68
	TuringLexerASSIGNMENT       = 69
	TuringLexerPLUS             = 70
	TuringLexerMINUS            = 71
	TuringLexerMULTIPLY         = 72
	TuringLexerDIVIDE           = 73
	TuringLexerDIV              = 74
	TuringLexerMOD              = 75
	TuringLexerREM              = 76
	TuringLexerEXP              = 77
	TuringLexerLESSTHAN         = 78
	TuringLexerGREATERTHAN      = 79
	TuringLexerEQUAL            = 80
	TuringLexerLESSTHANEQUAL    = 81
	TuringLexerGREATERTHANEQUAL = 82
	TuringLexerNOTEQUAL         = 83
	TuringLexerAND              = 84
	TuringLexerOR               = 85
	TuringLexerIMPLIES          = 86
	TuringLexerIN               = 87
	TuringLexerSHR              = 88
	TuringLexerSHL              = 89
	TuringLexerXOR              = 90
	TuringLexerPLUSEQUALS       = 91
	TuringLexerMINUSEQUALS      = 92
	TuringLexerMULTIPLYEQUALS   = 93
	TuringLexerDIVIDEEQUALS     = 94
	TuringLexerDIVEQUALS        = 95
	TuringLexerMODEQUALS        = 96
	TuringLexerANDEQUALS        = 97
	TuringLexerOREQUALS         = 98
	TuringLexerIMPLIESEQUALS    = 99
	TuringLexerSHREQUALS        = 100
	TuringLexerSHLEQUALS        = 101
	TuringLexerXOREQUALS        = 102
	TuringLexerIDENTIFIER       = 103
	TuringLexerWHITESPACE       = 104
	TuringLexerBLOCK_COMMENT    = 105
	TuringLexerLINE_COMMENT     = 106
)