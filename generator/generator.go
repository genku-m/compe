package generator

import (
	"fmt"
	"os"
)

type Generator interface {
	Generator() error
}
type generator struct {
	competition     string
	competitionType string
	title           string
}

func NewGenerator(c, ctype, title string) (*generator, error) {
	return &generator{
		competition:     c,
		competitionType: ctype,
		title:           title,
	}, nil
}

func (g *generator) Generate() error {
	if err := os.Mkdir(g.title, 0777); err != nil {
		return err
	}

	switch g.competition {
	case "atcoder":
		if err := g.atCoderContest(); err != nil {
			return err
		}
	default: // create a file
		if err := g.createFile(g.title); err != nil {
			return err
		}
	}
	return nil
}

func (g *generator) createFile(title string) error {
	fp, err := os.Create(title + ".cpp")
	if err != nil {
		return err
	}
	defer fp.Close()

	fp.Write(([]byte)(template()))
	if err := os.Rename(title+".cpp", g.title+"/"+title+".cpp"); err != nil {
		fmt.Println(err)
	}
	return nil
}

//what can't be done
func template() string {
	return `#include <bits/stdc++.h>
#define rep(i, n) for(int i = 0; i < n; i++)
using namespace std;
using ll = long long;
using P  = pair<int, int>;

int main() {}
`
}
