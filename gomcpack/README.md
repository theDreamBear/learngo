# gomcpack
本模块从百度的gomcpack移植而来，by zhouzhaopeng 2019.01.15

**gomcpack** is an go language framework for internal communication with nshead protocol and mcpack serialization. Currently, three components inclueded, they are ***mcpack***, ***mcpacknpc*** and ***npc***. ***mcpack*** is an package implement mcpack protocol. It behaves as consistent as possible with libmcpack, which is the C\C++ implementation of mcpack protocol.


## mcpack ##
### Usage ###

#### Marshal ####

##### Example #####

    type T struct {
        A bool
        X string
        S int
        V int   
    }
    
    t := &T{A:false, X:"ksarch", S:1, V:0}
    b, err := Marshal(t)

A struct is required, there are a []byte and an error returned. if success, an nil error will be returned.

#### Marshal with tag ####

##### Example #####

    type T struct {
        A bool
        X string
        S int
        V int `mcpack:"dandan"`
    }
    
    t := &T{A:false, X:"ksarch", S:1, V:0}
    b, err := Marshal(t)
    
As an tag ***dandan*** was applied to V, the user will see something like this after unmarshal:

    TT{A:false, X:"ksarch", S:1, dandan:0}

#### Marshal with exception ####

##### Example #####

    type T struct {
        A bool
        X string
        S int
        V int `mcpack:"-"`   
    }
    
    t := &T{A:false, X:"ksarch", S:1, V:0}
    b, err := Marshal(t)

When an tag ***-*** was applied to V, the variable V will not be serialized. The user will get the result as below after unmarshal:

    TT{A:false, X:"ksarch", S:1}

#### Unmarshal ####

##### Example #####

    type T struct {
        A bool
        X string
        S int
        V int    
    }
    
    t := &T{A:false, X:"ksarch", S:1, V:0}
    b, err := Marshal(t)

    tt := &TT{}
    err := Unmarshal(b, tt)

Two parameters are required, a []byte and a struct reference. An nil error will be returned when success.

#### Unmarshal with interface ####

    type III struct {
        S string
    }
    
    type II struct {
        B interface{}
        Name string   
    }
    
    type I struct {
        A *II    
    }

    t := &I{A:&II{B:&III{S: "V"}, Name: "SSG69"}}
    b, err := Marshal(t)

    tt := &I{A:&II{}}
    err = Unmarshal(b, tt)

Then tt would be seen like this after unmarhal:

    I{A: &II{B: (map[stirng]interface){"S": (string)"V"}, Name: "SSG69"}}

#### Unmarshal with unmarshaler ####

##### Example #####

    type T struct {
        data []byte    
    }

    type L struct {
        Name string
    }

    func (t *T) UnmarshalMCPACK(data []byte) error {
        if len(data) > 0 {
            t.data = data
            return nil    
        }
        return fmt.Errorf("invaildate data len: %d\n", len(data))
    }

    l := &L{Name: "SSG3000"}
    b, err := Marshal(l)

    t := &T{}
    err = Unmarshal(b, t)

Then t.data will get all bytes in b.

