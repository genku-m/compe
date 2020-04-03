package generator

import "errors"

func (g *generator) atCoderContest() error {
	switch g.competitionType {
	case "abc":
		if err := g.createFileForABC(); err != nil {
			return err
		}
	case "agc":
		if err := g.createFileForAGC(); err != nil {
			return err
		}
	default:
		return errors.New("unknown contest")
	}

	return nil
}

// TODO Get question from atcoder website
func (g *generator) createFileForABC() error {
	tasks := []string{"a", "b", "c", "d", "e", "f"}
	for _, v := range tasks {
		if err := g.createFile(v); err != nil {
			return err
		}
	}
	return nil
}

// TODO Get question from atcoder website
func (g *generator) createFileForAGC() error {
	tasks := []string{"a", "b", "c", "d", "e", "f"}
	for _, v := range tasks {
		if err := g.createFile(v); err != nil {
			return err
		}
	}
	return nil
}
