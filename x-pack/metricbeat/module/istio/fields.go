// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package istio

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "istio", asset.ModuleFieldsPri, AssetIstio); err != nil {
		panic(err)
	}
}

// AssetIstio returns asset data.
// This is the base64 encoded gzipped contents of module/istio.
func AssetIstio() string {
	return "eJzUW81y47gRvvspuvaUpDx8AB9SldpJKnPYrcnOHFK5eCGwZWINAgzQkKx5+lTjh6JpirYl0nZ08Egkge/rH3Q3GpxPcI+HG1CelL0CIEUab+Cn+PunK4AavXSqI2XNDfz1CgDSs/CLrYPGKwCHGoXHG9ggiSuArUJd+5v46CcwosXj9PyhQ4c3cOds6PKVCQz+/B5H/Q7SGhLKePAkiK9JD9QIgj06BIeihq2zLXwZgAxJDIm06Jv+4hSXGT78+XmCi0MtCGsgC9RgohGBwKPbKYmDGcbKKp8x30fKM56EeTTNkTvhA41uzNDnz/cGoXO2RWow+OnZC/QfdjOJeo+HvXX1RcDjuQumw/8G9OQngbU1d69EtSR0Pyk0wtQaa9gcQJhsq87Zh8Mcl6oOTjBA1fpqE+Q9UvWXSX528wfKsUHSxdtzJfgtkYBCAhrlyd450ULiwkaEVmmtPEprav9iWXxoX6rmrXWtoJuexVky+H74NfjQgt32F84XQtpgxiq/UNdDnia0G3RMddIzx8y8+oHV5kD4jp7yTf3Ap17yUtqvcIrz+IwBZrisYNsxm8cQJ4j5zhqP/3/Gfcr7Ha37lMxHMW9nHaFbLtX9lmcEVaMhtVXoY21QkDiapN/J4+ELgfLgkbiKqNGTMinWq20exPdjmSNiWYFumLxAmBq8DU7i5AipFRoajqgmNZGmqPbW3Wsr6oqvLlkAKD/WCCOwNjL5ggz7Rskmln7O6vRoJvdi4r4TJ4qmRdlHmGLQkRizZDunjFSd0Oty7BAd9FiFKDmx3SpZuAwccN+gSYNEoIanktkVPQSP8zKJrltXmqxg0XWwER5rsCb+0GKD+hwr7ND5p+XMwqwzyLP8JokOwsF7rMxhNJrX6Emi77ESz6P9IdbkkNDFC3M42eqrc6j0+SU6aZ7nJXiP1TrvSc9Szk2AqrF+utRZhPeQZEYERixll9po5D0c73qVkbZV5q4vP14sw/pRZ0oORn0dx7cOOBOsn/GQovrOWbLSrh1vMsp83fm3r1+OT6otf9+pGutrsNSg2yvfjwTruDgzKKPMZdS0lfrKX9p62ipntHUmpCw4wDhPRI0jUgxSHjqHnktia/SBA9Q/v3//Ci2SU9JPS3EUt/Iog1N0qDqrlTysXPHkVTAK+gl63qBtoCD0LWmfkkeq/3O6iA+Ie9ZW2wZT5o2yYdxQjLYRJ9ISWQjm3ti9SSCjUbnW8crw36y4wl4KYyzBJnpoh07znS7EfurRCn3zVj082qSt2L1lpKXatzxj1cquKkYS8n6x7iZHonGDChgAHEpUu9TqfG77JK3ZqrtK1KIjdJUyW1uhc9bxFp1vrUI3IQCauEdHhzXUwXFi6pyV6D1/zQ6eqQFTy3RLC/DFEqVr66g+LYAMF3uZzFoG5zjKZDazTImc2gTCVfn1IGcwTK1z9wZqzEiXkMzeuxNa1VP96tU8eINSBI9FBjgyAIcUnME6DZsVo5zKfLBVWGi9YgX2kqzvNgXqHL9xQX80ZTOlVyh6KEErSDbrpJhMqlbp0CYdxAoPnMc74TyKjZ7PMpHo+t7AMOd4AmHbcUnwwbyh0HqFR/SSrK/sAnWOwoPxXI1tFdaViAX2bQ6eqxBOENlzt0LpqHXkArDPOxtkzQcjdkLp0w5dK9/xSkN3OyiO/W2H7nYnnEI63JKlE92kSwX5JRappTAaEoglX/GXTCSCTotRcmYRv5LaehxvXZYh3av4+AZHQiuen5wCyAnj1WkXL5xrga01S7p28dzHHrNDSFDgbCBlkqcLuFM77Os+QLNTzpoWzYmmSqHNfhdcPEZU+i00nR1dOky1yHna7mlHm71VTNw3Sic3Ycq9XOeJ0Du5wf1bebjBvT4k3Z/t6D1vh+ENF6fDT+Hs1bkkyXwaEXuEaSFOohQGOeqtxGA8e0HtlLb0Jh2KiLRQh+KhnjuJV4bw7oktn1HYzzwda0taUwLqw+dv5SD62EvzsSn9laWZdiUm1wXfnNgXX8aOKbXovbhDDx4NXYPwsEet+d8cimKIjpWY42fi1+HAJ9OzRLr21+D4j6x97KZhfaKxWCSsSLVYeZRLOW18840n5VSV36pKqgYS9+jZoxj4tWT5Zr/XmiQ7+f7LM3S/9Ls31h8K2YDUwRO6aBS7BS08RcIV/AedHe73mPXctrrXsbSG8IFycb+cS/06TmB/YsXbQP7PoIwiJSjW9azuwuEkU2bgjNCLs0wucQz4BQj+/flbIa5MCi6na3atJFVaeUITW2sbG8x0QrpQlQWMNZdhoOCe7tKP+NlAcWTVEHWV3aGrckSqSE6fhy5Ie690LYWrgdGP5GGvqOlLzf4hksdnLpCPv63i20PJ4onJSKAYOdOOkG9/g846er2lSI4MxQK9maWGRjhlqEfWvFy+N3DEGame97r0ypi0Zvcub0B/5xxmEOt0mpFTWE5efzc7e3jcFfHXz78j/RJBl3tX8iIJ8pvSnFBeKEauCxfPHGXe0vWxc6lipxxVazHhycNrGT227pIvn15kX1nK0WjhaerGVmo6RpxTa321derYbmNOzT07NHVnFQcEsdF4DZ31Xm30AZSJpxgn4lwq0ia5nfe/U/pThn/86/Ov06A8/YKIyULpVdpHdspds/8FAAD//9kkN0U="
}
