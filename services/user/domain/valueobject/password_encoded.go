package uservo

type PasswordEncoded string

func (pe PasswordEncoded) Equals(other PasswordEncoded) bool {
	return pe.String() == other.String()
}

func (pe PasswordEncoded) String() string {
	return string(pe)
}

func NewPasswordEncoded(value string) (encodedPassword PasswordEncoded) {
	return PasswordEncoded(value)
}
