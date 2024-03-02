package spotify

import (
   "fmt"
   "os"
   "testing"
)

func TestSolution(t *testing.T) {
   username := os.Getenv("spotify_username")
   password := os.Getenv("spotify_password")
   var req login_request
   if !req.New(username, password) {
      t.Fatal("Getenv")
   }
   res, err := req.login()
   if err != nil {
      t.Fatal(err)
   }
   suffix, iterations, err := res.solve_hash_cash_challenge()
   if err != nil {
      t.Fatal(err)
   }
   solution, err := res.challenge_solution(
      username, password, suffix, iterations,
   )
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%#v\n", solution)
}
