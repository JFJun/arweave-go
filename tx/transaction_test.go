package tx

import (
	"fmt"
	"github.com/JFJun/arweave-go/utils"
	"github.com/JFJun/arweave-go/wallet"
	"testing"
)

func TestNewTransactionV2(t *testing.T) {

	/*
	{
	  kty: 'RSA',
	  n:   'y244eCb3M3dEx-NqjlLFWmF2KDFyO4qr3J5wSht0NzToFRRocb60CvZg0G5rR3xrMzJPg4PgP86Z-K8BqMRN_5uXlvkG5ur0Uczzen4Icb0ym_hQeZqxMMLbinY86nr9Bl0kf3ddnAZNJlSqKStBIfJEQoNWeAEHAvEUort9qDy0g5CYn1CrWMAcdE13hUaJ0a6YAGNq1ACxc3D1UCeUn14xze4CAItXJiTK18rIN8L1ZI0o6KqpfMMidrQxgmuQR18UpFw3v-DSicA2DLfVxZm8WgWtEq8G-0M2w60ToGOp1fhDQeUwEhMuVBnG6Fm7TgtBQXnMEnMI48RadvtT5udJG-DcdPXxePsFkzPst_2jXofUT1ENfIa8Nr_A_IHbiFFWUccqTL4sUnOcA38wZ5loe2SJ8n5Iww1LRPwi8o9NVFKNggGKY5JV93r_Cd6Kl3zjgMaDg7CaSTp2DR9u_4T9VR0JuGJsj1rKizUkWqvIcHLqdqBVg0eTlUYkj9GcNsL2gO0f_wRWyjgCue6LEsZo8axgPpxvAchdRRbIbrjVnE0UwVnL8gFNevtr16ADXJPrWxqWMdNFVYIQgp0-WR70IJJfL2rmSsHe3AcRxqVkB4qJ0VWirMmcmcCWX6OiX_t9_Q3FxPavtfNegCHBD2_L7Sh30udJal7qScWxlPM',
	  e:   'AQAB',
	  d:   'ncUWwICf2WdXjZrfShABhQ3vQE7Q5EKw9UJfnena4PIESEzyj7yyD4pzHVN1mhRxY1xSdJHzeij2GvtHUycjwuhd0bXC15ZC86n67vYOAIAWXwZfC2CgqUuGTsVSixvlXmkgZEcj71ynuuGQsqscMYKd0-PeTOWnFW35bhfw5WhnK14aL_y1VvhvWPe3iUJiop9AvZRdBpyfAwQJDSgVSOAZyjcbkGaAqVGBi4h2bWCjRl9_6apqFhuUf_FSMG7Rj7i63pmylvp67mCGyqhMf5qyjDK4vrPFu1quf2WcwhhpYEvRhaqE9I-xr4HqVMs6y9VIv8oBLhHPRC5FZt-WeiwqrOwx8Orazc1smKXp8CyPwefIhUG861C-BDzNgmMDiU8gFPgikOGbAQOxbG7U_jTjJUUzjxXuLUX9Uif1BjOpvdImxU8fGSv30HSAwsLjWH29Ovm9FGMwOWs83-AlPK7hZiWyTn0t5TzafsesTEm-CNt4cnBDXDqtxxyUhC_ksrKMKFEANsxgNASWT5kuDZo-gxon2KPU-jiXN6xBp6l81GIbkJnMA5AyutrXBcMDf0JShGhHQ-2Ih66R2iP-9TJaISjjkxG2iPY2WZ0US-KxG7E_zRhdrixhg-M1dMq9hneCLan4hzakmG-R9bhrEvywKpE1hTjTm0dKyqD3PEE',
	  p:   '8ldDyt2XdyhpCs-PwXOmh_egoxY82vgwfmX3oLFePfFrL7NtssEym0rwf7FKl3KCTj2pSjixmh-zsUcFeAPLBTZtV9DBFHhT-K0jhYGuvnT7H1Q7tjn2tEW4jl3cVxFovNsS4_ue4JcBmQJ8kkExT9Ek3W1BsBvoL0WYM_oMD5mQ6YUtARxG9V5dHdFg_MwZIARzGNcAYqd-CyM-U55cEpsXK6GJf-2UlFzJ64CR0Bux8JrsVrTwm2eMcuNP6gLpoqeiRNwOLr5tsHKsBj4_-LauQaP8xmaynoq3_JHzO9xZmFzih5Sx-KdHaBAy8eY8QoZ5IkwBtuQHbzk6Cd0Y0w',
	  q:   '1uWE0zMa8Fv0QB8D-8fBijiOJyyfTlkLmEYapDejioY4LNNrQ3NYvdHZ9BsnjM6jMAhneWk4vpSOnufqfYM83xLpJ-SxOOF8T9d-7tR_3qmKpqdOarseEPZr1iodlnYiN4zTLGKwoXUozsrnnb1NCfoiB1eNjOW8xSOD4uYJdwJwOPRIU6Bvy0hEDzItdH1_0N8XQVj8z41pOmqfx515VjK2vMRPdhQcFBOLmK0KlJdvic3xgjfayHE7xkXUNR4vusHXOHCmgSKd8H27L17tQw7G0OA844Ov20e2dOEH1jORI0vJ6UHn6FdCvTzWfDtzqs9jBfv_Icun7zkOv2v_YQ',
	  dp:  '3s8qzjbkDrNBuKXTVxqcrwAWWjuU4gI0m9HmVYEd07fGQ8olfqcwfyTsb4qyuQGYGpWIs2vipoClNsXEmm5qV7WN3KJbExX6pnWRZiswXX2ycUCB0e4DDAXaRFK5xfeEo7aL7L6oeXd9CZAGtnVoACmbSlYsgd12oRfMc2Vfd2xKltlR4mZ8OxZyaHrcQDDqnMxagikS-qFiOp1BraG4p4aeTAguIkduRuX05CiGZ58-_6eAyct8OHWA5RRnohDhfCHCCKuGAqYktYkI7oiHSoWzOziprEQ5z5hcxsPtrUjPu00THXKyGpq6BUx3en6jQsreTpExNlT7HOCxouX77w',
	  dq:  'LeIZpyKNSO7JSvuUVSkEiOekanbbmNXoZN3rCE8c6xHXYrLNfzxcoULyP22g1y6zpyjUNxchV4fOn90ysxPXZmXaRHtO3689ZKrFXmce2kLm6MhmNOG4_LQwuCyL3rlgeu3llN-JiR_D08t9ArANh-jRZTHrdj8DPDBmE4c8VB3AKlj2LGzgEc5fQcNb8zhslwNbKse-Su6Nnjxv2yYjZyUy4UElGjgW-q5lqvafgLSIyyqeu0gDvbJ3qdD1C1ZjqIOE2XcsBfyHLncaBXZGQIAiMSz_szBF1xMLwQ7fJnwiLCAdf2FkC7a9-DvKLAJhowHeycNRaTzboXdUT8jWYQ',
	  qi:  'dwvprMEgrqTpZmUuX8sC7NzeHpZAWybikUnO9y-9n_rDIINPhxs6HOM8NB4J2j5VGINJaQKrCWIsvxF0hTC5LR6R_TIuyTdZjHJ21haG73yUzLhSalvB4FnBBRA8MA5XA2EhJjRg9fNym5GAoqVhvKqKdGc9aWHzNGN3DftL4IfaO-F0empeBdFAmezegpHp6EfUi3sAqGsHrqmeF05Grc-ThkMXpA7Ie96TO7e7_6pgrkfPLHHISuY3e5PU_W3kAsktZHimHL_ch6Aj84wQ4EUe9HPTOIwgFkCbh2S1PRzpszIawSTWAN2Yu1xFpzWWNtAjToNSDAPGotH-BM6O2Q'
	}
	Transaction {
	  format: 2,
	  id: '',
	  last_tx: 'exMh68dK_MbcAEB-NY2tMNAO7w4-jaJSqV1RxJY4hzmtTVCtQujJP9zRdiwcNmIH',
	  owner: 'y244eCb3M3dEx-NqjlLFWmF2KDFyO4qr3J5wSht0NzToFRRocb60CvZg0G5rR3xrMzJPg4PgP86Z-K8BqMRN_5uXlvkG5ur0Uczzen4Icb0ym_hQeZqxMMLbinY86nr9Bl0kf3ddnAZNJlSqKStBIfJEQoNWeAEHAvEUort9qDy0g5CYn1CrWMAcdE13hUaJ0a6YAGNq1ACxc3D1UCeUn14xze4CAItXJiTK18rIN8L1ZI0o6KqpfMMidrQxgmuQR18UpFw3v-DSicA2DLfVxZm8WgWtEq8G-0M2w60ToGOp1fhDQeUwEhMuVBnG6Fm7TgtBQXnMEnMI48RadvtT5udJG-DcdPXxePsFkzPst_2jXofUT1ENfIa8Nr_A_IHbiFFWUccqTL4sUnOcA38wZ5loe2SJ8n5Iww1LRPwi8o9NVFKNggGKY5JV93r_Cd6Kl3zjgMaDg7CaSTp2DR9u_4T9VR0JuGJsj1rKizUkWqvIcHLqdqBVg0eTlUYkj9GcNsL2gO0f_wRWyjgCue6LEsZo8axgPpxvAchdRRbIbrjVnE0UwVnL8gFNevtr16ADXJPrWxqWMdNFVYIQgp0-WR70IJJfL2rmSsHe3AcRxqVkB4qJ0VWirMmcmcCWX6OiX_t9_Q3FxPavtfNegCHBD2_L7Sh30udJal7qScWxlPM',
	  tags: [],
	  target: 'GRQ7swQO1AMyFgnuAPI7AvGQlW3lzuQuwlJbIpWV7xk',
	  quantity: '1500000000000',
	  data: '',
	  data_size: '0',
	  data_root: '',
	  data_tree: [],
	  reward: '6475681',
	  signature: ''
	}
	*/
	jwk:=`{
	  "kty": "RSA",
	  "n": 	 "y244eCb3M3dEx-NqjlLFWmF2KDFyO4qr3J5wSht0NzToFRRocb60CvZg0G5rR3xrMzJPg4PgP86Z-K8BqMRN_5uXlvkG5ur0Uczzen4Icb0ym_hQeZqxMMLbinY86nr9Bl0kf3ddnAZNJlSqKStBIfJEQoNWeAEHAvEUort9qDy0g5CYn1CrWMAcdE13hUaJ0a6YAGNq1ACxc3D1UCeUn14xze4CAItXJiTK18rIN8L1ZI0o6KqpfMMidrQxgmuQR18UpFw3v-DSicA2DLfVxZm8WgWtEq8G-0M2w60ToGOp1fhDQeUwEhMuVBnG6Fm7TgtBQXnMEnMI48RadvtT5udJG-DcdPXxePsFkzPst_2jXofUT1ENfIa8Nr_A_IHbiFFWUccqTL4sUnOcA38wZ5loe2SJ8n5Iww1LRPwi8o9NVFKNggGKY5JV93r_Cd6Kl3zjgMaDg7CaSTp2DR9u_4T9VR0JuGJsj1rKizUkWqvIcHLqdqBVg0eTlUYkj9GcNsL2gO0f_wRWyjgCue6LEsZo8axgPpxvAchdRRbIbrjVnE0UwVnL8gFNevtr16ADXJPrWxqWMdNFVYIQgp0-WR70IJJfL2rmSsHe3AcRxqVkB4qJ0VWirMmcmcCWX6OiX_t9_Q3FxPavtfNegCHBD2_L7Sh30udJal7qScWxlPM",
	  "e":   "AQAB",
	  "d":   "ncUWwICf2WdXjZrfShABhQ3vQE7Q5EKw9UJfnena4PIESEzyj7yyD4pzHVN1mhRxY1xSdJHzeij2GvtHUycjwuhd0bXC15ZC86n67vYOAIAWXwZfC2CgqUuGTsVSixvlXmkgZEcj71ynuuGQsqscMYKd0-PeTOWnFW35bhfw5WhnK14aL_y1VvhvWPe3iUJiop9AvZRdBpyfAwQJDSgVSOAZyjcbkGaAqVGBi4h2bWCjRl9_6apqFhuUf_FSMG7Rj7i63pmylvp67mCGyqhMf5qyjDK4vrPFu1quf2WcwhhpYEvRhaqE9I-xr4HqVMs6y9VIv8oBLhHPRC5FZt-WeiwqrOwx8Orazc1smKXp8CyPwefIhUG861C-BDzNgmMDiU8gFPgikOGbAQOxbG7U_jTjJUUzjxXuLUX9Uif1BjOpvdImxU8fGSv30HSAwsLjWH29Ovm9FGMwOWs83-AlPK7hZiWyTn0t5TzafsesTEm-CNt4cnBDXDqtxxyUhC_ksrKMKFEANsxgNASWT5kuDZo-gxon2KPU-jiXN6xBp6l81GIbkJnMA5AyutrXBcMDf0JShGhHQ-2Ih66R2iP-9TJaISjjkxG2iPY2WZ0US-KxG7E_zRhdrixhg-M1dMq9hneCLan4hzakmG-R9bhrEvywKpE1hTjTm0dKyqD3PEE",
	  "p":   "8ldDyt2XdyhpCs-PwXOmh_egoxY82vgwfmX3oLFePfFrL7NtssEym0rwf7FKl3KCTj2pSjixmh-zsUcFeAPLBTZtV9DBFHhT-K0jhYGuvnT7H1Q7tjn2tEW4jl3cVxFovNsS4_ue4JcBmQJ8kkExT9Ek3W1BsBvoL0WYM_oMD5mQ6YUtARxG9V5dHdFg_MwZIARzGNcAYqd-CyM-U55cEpsXK6GJf-2UlFzJ64CR0Bux8JrsVrTwm2eMcuNP6gLpoqeiRNwOLr5tsHKsBj4_-LauQaP8xmaynoq3_JHzO9xZmFzih5Sx-KdHaBAy8eY8QoZ5IkwBtuQHbzk6Cd0Y0w",
	  "q":   "1uWE0zMa8Fv0QB8D-8fBijiOJyyfTlkLmEYapDejioY4LNNrQ3NYvdHZ9BsnjM6jMAhneWk4vpSOnufqfYM83xLpJ-SxOOF8T9d-7tR_3qmKpqdOarseEPZr1iodlnYiN4zTLGKwoXUozsrnnb1NCfoiB1eNjOW8xSOD4uYJdwJwOPRIU6Bvy0hEDzItdH1_0N8XQVj8z41pOmqfx515VjK2vMRPdhQcFBOLmK0KlJdvic3xgjfayHE7xkXUNR4vusHXOHCmgSKd8H27L17tQw7G0OA844Ov20e2dOEH1jORI0vJ6UHn6FdCvTzWfDtzqs9jBfv_Icun7zkOv2v_YQ",
	  "dp":  "3s8qzjbkDrNBuKXTVxqcrwAWWjuU4gI0m9HmVYEd07fGQ8olfqcwfyTsb4qyuQGYGpWIs2vipoClNsXEmm5qV7WN3KJbExX6pnWRZiswXX2ycUCB0e4DDAXaRFK5xfeEo7aL7L6oeXd9CZAGtnVoACmbSlYsgd12oRfMc2Vfd2xKltlR4mZ8OxZyaHrcQDDqnMxagikS-qFiOp1BraG4p4aeTAguIkduRuX05CiGZ58-_6eAyct8OHWA5RRnohDhfCHCCKuGAqYktYkI7oiHSoWzOziprEQ5z5hcxsPtrUjPu00THXKyGpq6BUx3en6jQsreTpExNlT7HOCxouX77w",
	  "dq":  "LeIZpyKNSO7JSvuUVSkEiOekanbbmNXoZN3rCE8c6xHXYrLNfzxcoULyP22g1y6zpyjUNxchV4fOn90ysxPXZmXaRHtO3689ZKrFXmce2kLm6MhmNOG4_LQwuCyL3rlgeu3llN-JiR_D08t9ArANh-jRZTHrdj8DPDBmE4c8VB3AKlj2LGzgEc5fQcNb8zhslwNbKse-Su6Nnjxv2yYjZyUy4UElGjgW-q5lqvafgLSIyyqeu0gDvbJ3qdD1C1ZjqIOE2XcsBfyHLncaBXZGQIAiMSz_szBF1xMLwQ7fJnwiLCAdf2FkC7a9-DvKLAJhowHeycNRaTzboXdUT8jWYQ",
	  "qi":  "dwvprMEgrqTpZmUuX8sC7NzeHpZAWybikUnO9y-9n_rDIINPhxs6HOM8NB4J2j5VGINJaQKrCWIsvxF0hTC5LR6R_TIuyTdZjHJ21haG73yUzLhSalvB4FnBBRA8MA5XA2EhJjRg9fNym5GAoqVhvKqKdGc9aWHzNGN3DftL4IfaO-F0empeBdFAmezegpHp6EfUi3sAqGsHrqmeF05Grc-ThkMXpA7Ie96TO7e7_6pgrkfPLHHISuY3e5PU_W3kAsktZHimHL_ch6Aj84wQ4EUe9HPTOIwgFkCbh2S1PRzpszIawSTWAN2Yu1xFpzWWNtAjToNSDAPGotH-BM6O2Q"
	}`

	/*
	*/
	w:=wallet.NewWallet()

	err:=w.LoadKey([]byte(jwk))
	if err != nil {
		panic(err)
	}

	v2:=NewTransactionV2("exMh68dK_MbcAEB-NY2tMNAO7w4-jaJSqV1RxJY4hzmtTVCtQujJP9zRdiwcNmIH",w.PubKeyModulus(),"1500000000000","GRQ7swQO1AMyFgnuAPI7AvGQlW3lzuQuwlJbIpWV7xk",nil,"6475681")
	//v2.AddTag("test-tag-1","test-value-1")
	//v2.AddTag("test-tag-2","test-value-2")
	//v2.AddTag("test-tag-3","test-value-3")
	d,_:=v2.MarshalJSON()
	fmt.Println(string(d))
	tx,err1:=v2.Sign(w)
	fmt.Println(utils.EncodeToBase64(tx.signature))
	if err1 != nil {
		panic(err1)
	}

}