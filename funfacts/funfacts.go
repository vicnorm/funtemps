package funfacts

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra
  Sett inn alle Funfucts i en struktur
  type FunFacts struct {
      Sun []string
      Luna []string
      Terra []string
  }
*/

// Funfacts i en struktur
type FunFacts struct {
  Sun []string
  Luna []string
  Terra []string 
}

// funfacts-funksjon
GetFunFacts(about string) []string 

  var funfact FunFacts

      funfact.Sun = []string{"Temperatur i Solens kjerne", "Temperatur på ytre lag av Solen"}
      funfact.Luna = []string{"Temperatur på Månens overflate om natten", "Temperatur på Månens overflate om dagen"}
      funfact.Terra = []string{"Høyeste temperatur målt på Jordens overflate", "Laveste temperatur målt på Jordens overflate", "Temperatur i Jordens indre kjerne"}

      if about == "sun" {
          return funfact.Sun
      } else if about == "luna" {
          return funfact.Luna
      } else if about == "terra" {
          return funfact.Terra
      }
      return nil
