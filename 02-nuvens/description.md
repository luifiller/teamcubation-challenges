**Avaliação: Saltar nas Nuvens**

---

### Descrição

Você recebe um slice de nuvens representado por inteiros, onde:

- `0` = nuvem segura
- `1` = nuvem perigosa (não pode pisar)

Seu objetivo é ir da primeira à última nuvem com o **mínimo número de saltos** possível. De uma nuvem segura você pode saltar **1** ou **2** posições para frente, mas somente se a nuvem de destino também for segura.

Implemente em Go a função:

```go
// saltarEmNuvens recebe um slice c e retorna
// o número mínimo de saltos para alcançar o final.
// Assume-se que c[0] = 0 e c[len(c)-1] = 0.
func saltarEmNuvens(c []int) int
```

E um programa `main` que:

1. Lê da entrada padrão:

   - Um inteiro `n` (2 ≤ n ≤ 100): quantidade de nuvens.
   - Uma linha com `n` inteiros (`0` ou `1`), separados por espaço.

2. Chama `saltarEmNuvens(c)`.
3. Imprime o resultado (um inteiro) na saída padrão.

---

### Formato de entrada

```
n
c[0] c[1] c[2] … c[n-1]
```

- `n`: número de nuvens.
- `c[i]`: estado da nuvem i (0 = segura, 1 = perigosa).

### Formato de saída

Um único inteiro: a **quantidade mínima de saltos** para ir da nuvem 0 até a nuvem n-1.

---

### Exemplo

**Entrada**

```
7
0 1 0 0 0 1 0
```

**Saída**

```
3
```

**Explicação**
Para `c = [0,1,0,0,0,1,0]`, a sequência mínima de saltos é:

- 0 → 2
- 2 → 4
- 4 → 6
  Total = 3 saltos.

---

### Restrições e casos de teste

- Inicia e termina em nuvens seguras (`c[0] = c[n-1] = 0`).
- Garante-se que exista ao menos um caminho até o final.
- Teste sua solução com casos extremos, por exemplo:

  - `n = 2`, `c = [0,0]` → resposta `1`
  - Todos zeros: `c = [0,0,0,0,…]`
  - Padrões alternados de 0 e 1.

---

### Critérios de avaliação

- **Corretude**: retorna o número mínimo de saltos em todos os cenários.
- **Clareza**: código legível, nomes descritivos e comentários, se necessário.
- **Eficiência**: complexidade O(n) e uso extra de memória O(1).
- **Boas práticas em Go**:

  - Tratamento de erros de leitura.
  - Código formatado com `gofmt`.
  - (Opcional) testes unitários e benchmarks.

---

Boa implementação e bons saltos!
