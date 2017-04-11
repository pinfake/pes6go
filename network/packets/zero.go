package packets;

type Zero struct {
    Zero uint32;
}

func NewZero() Zero {
    z := Zero{0};
    return z;
}