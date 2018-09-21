# ker

API wrapper for Hostker written in pure Go

## Documentation

See https://godoc.org/github.com/violarulan/ker

## Example

    package main
    
    import (
        "github.com/violarulan/ker"
        "fmt"
    )
    
    func main(){
        var resp ker.ApiResponse
        user := ker.User{Email: "user@domain.tld", AccessToken: "d0113297439b0a467fbd30dbd4ad369a1e4e0189807b63575aff94c038600005"}
        resp = user.Balance()
        fmt.Println(resp.Balance)
    }

## Contribute

Feel free to submit issues and PRs.
