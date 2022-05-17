# Tipos de Dados - Capacidades

Embora seja um assunto não muito abordado, pelo menos no meu ponto de vista, entender todos os tipos de dados que uma linguagem oferece, ajuda muito no momento de criar um programa, principalmente do ponto de vista de performance e eficiência no consumo de recursos.

Por isso, no post de hoje vamos ver todos os tipos de dados básicos que temos na linguagem e suas capacidades.

## **Bool**

O tipo booleano, é um dos tipos mais básicos e comum no mundo da tecnologia.

Embora em algumas linguagens ele possa ser definido com **0** e **1**, em Go os únicos valores aceitos são **true** ou **false**.

## String

Conjunto de todos os caracteres de 8-bit, ou seja, qualquer tipo de texto.

Uma string pode ser vazia, mas não pode ser nula (_nil_).

## Inteiros

Dentro do mundo dos inteiros temos vários tipos que nos ajudam a dimensionar melhor a utilização de recurso dentro de um programa.

| int       | capacidade: 32bits em sistemas 32bits e 64bits em sistemas 64bits, mas não é um alias para int32 ou int64 |
| --------- | --------------------------------------------------------------------------------------------------------- |
| int8      | capacidade: -128 ~ 127                                                                                    |
| int16     | capacidade: -32.768 ~ 32.767                                                                              |
| int32     | capacidade: -2.147.483.648 ~ 2.147.483.647                                                                |
| int64     | capacidade: -9.223.372.036.854.775.808 ~ 9.223.372.036.854.775.807                                        |
| uint      | capacidade: 32bits em sistemas 32bits e 64bits em sistemas 64bits, mas não é um alias para uint32         |
| ou uint64 |
| uint8     | capacidade: 0 ~ 255                                                                                       |
| uint16    | capacidade: 0 ~ 65.535                                                                                    |
| uint32    | capacidade: 0 ~ 4.294.967.295                                                                             |
| uint64    | capacidade: 0 ~ 18.446.744.073.709.551.615                                                                |
| uintptr   | capcidade: grande o suficiente para guardar o padrão de bits de qualquer ponteiro                         |

A recomendação para esse tipo de dado é que, na dúvida sempre use **int**, mas caso você saiba que uma determinada variável ou atributo da sua struct tem um certo limite no seu tamanho, então use o tipo adequado para isso.

Por exemplo, em uma struct ou variável onde o objetivo é guardar a idade de uma pessoa, não há necessidade de um tipo que tenha números negativos ou maior do que 255, ou seja, nesse caso faria mais sentido usar o tipo **uint8** ao invés de **int.**

Um ponto importante aqui é que, sempre que iniciamos uma variável do tipo inteiro sem definir seu tipo, o compilador irá tratar essa variável como **int**.

## Byte

Esse tipo de dado é um alias para **uint8**, porém é utilizado para deixar claro que o dado esperado é um byte ao invés de um inteiro de 8-bits unsigned.

## Rune

Assim como byte, rune também é um alias, só que para **int32**. Esse tipo de dado é usado para distinguir caracteres de inteiros.

## Decimais

Apenas dois tipos são necessários para lidar com números do tipo float.

Um ponto de atenção aqui é que, quando iniciamos uma variável com um número decimal mas não definimos seu tipo, o compilador irá tratar ela como **float32**.

| float32 | capacidade: todos os pontos flutuantes IEEE-754 de 32 bits |
| ------- | ---------------------------------------------------------- |
| float64 | capacidade: todos os pontos flutuantes IEEE-754 de 64 bits |

## Números Complexos

Se você precisar guardar algum número com parte real e parte imaginária, como por exemplo **12.8i**, esse é o seu tipo de dado.

| complex64  | capacidade: todos os float32 reais e parte imaginária |
| ---------- | ----------------------------------------------------- |
| complex128 | capacidade: todos os float64 reais e parte imaginá    |

No caso de números complexos, quando não definimos o seu tipo ao iniciar uma nova variável, o compilador do Go irá trata-lo como sendo do tipo **complex128**.

Para finalizar, como foi citado acima, ao não declarar o tipo de dado, o compilador irá definir automágicamente um tipo para aquela variável.

Caso você queira iniciar sua variável com um tipo diferente do padrão, existem duas formas básicas para se fazer isso.

```go
i := uint8(10)
var f float64 = 10.59
```
