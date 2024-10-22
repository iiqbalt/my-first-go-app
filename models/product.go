type Product struct {
  Id int
  Nama string
}

func (p *Product) GetNama() string {
  return p.Nama
}
