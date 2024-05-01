# FLanguage

Benvenuto su FLanguage! Questo progetto è stato creato come esperimento per comprendere il funzionamento degli interpreti.

## Caratteristiche Principali

- **Funzioni**: Puoi definire e utilizzare funzioni per organizzare il tuo codice in moduli riutilizzabili e per gestire compiti specifici.Le funzioni non possono modificare l'ambiente esterno in cui sono definite.
  
- **Array**: Manipola collezioni di dati con facilità utilizzando array, consentendo di memorizzare e accedere a elementi in modo efficiente.
  
- **Hashtable**: Utilizza hashtable per implementare strutture dati chiave-valore, ideali per la gestione di associazioni tra dati e inline function.
  
- **Istruzioni Condizionali**: Controlla il flusso del programma utilizzando istruzioni condizionali if-else, consentendo di eseguire operazioni diverse in base a determinate condizioni.
  
- **Cicli While**: Itera attraverso i dati o esegui operazioni ripetute finché una condizione specifica è vera utilizzando i cicli while.
  
- **Funzioni Inline**: Definisci funzioni direttamente nel contesto del codice principale per migliorare la leggibilità e la modularità del codice.
  
- **Importazione di Codice Esterno**: Importa facilmente codice da altri file per organizzare il tuo progetto in moduli separati e riutilizzare il codice esistente.
  
- **Oggetti**:Vi è la possibilità di poter creare un hashtable con al interno delle innerfunction che potranno interaggire con la hashtable attraverso la parola chiave "this"
## Sintassi

- **let**: Utilizzata per dichiarare variabili.
  - Esempio: `let a = 2;`

- **import**: Utilizzata per l'importazione di file o moduli.
  - Esempio: `import("BinarySearch.txt");`

- **Ff**: Utilizzata per dichiarare una funzione.
  - Esempio:
    ```
    Ff getMatrix() {
         return [[2,4],[2,3,4]];
    }
    ```
- **@**: Utilizzate per definire funzioni direttamente nel contesto del codice principale, salvandole all'interno di variabili, hashtable o array.
  - Esempio:
    ```
	  let a=@(a,b){
		  ret a+b;
	  };
	  let b=a(2,1);
    ```
- **ret**: Utilizzata all'interno di una funzione per restituire un valore.
  - Esempio:
    ```
    Ff funzione(){
       ret this{"eta"};
    }
    ```

- **if/else**: Utilizzati per creare una struttura condizionale.
  - Esempio:
    ```
    if (4 < 2){
        a = a + 2;
    } else {
        a = a * 4;
    }
    ```

- **while**: Utilizzata per creare un ciclo che continua fintanto che la condizione specificata è vera.
  - Esempio:
    ```
    while (i < 5) {
       i = i + 1;
    }
    ```

- **newArray**:Funzione per creare un nuovo array con valori specificati.
  - Esempio: `let a = newArray(4, 0);`

- **len**: Funzione per determinare la lunghezza di un array o stringa.
  - Esempio: `let b = len(a);`

- **import()**: Utilizzata per l'importazione di file o moduli, anche se potrebbe variare a seconda del linguaggio di programmazione specifico.
  - Esempio: `import("nome_modulo");`

- **this{}**: Utilizzata per fare riferimento all'oggetto corrente all'interno di un contesto di programmazione orientato agli oggetti.
  - Esempio:
    ```
    let object={
               "nome":"luca",
               "eta":22,
               "compleanno":@(){
                    this{"eta"}=this{"eta"}+1;
                    ret this{"eta"};
               }
    };
    object{"compleanno"}();
    ```


- **string()**: Presumibilmente una funzione o un metodo per convertire un valore in una stringa.
  - Esempio: `a = a + string(2);`

- **getMatrix**: Presumibilmente una funzione per ottenere una matrice o una struttura dati simile.
  - Esempio: `let b = getMatrix()[0][1];`

## InnerFunction:

- **len**
  - Parametri: `a`
  - Funzione incorporata: `lenBuiltin`

- **newArray**
  - Parametri: `n`, `type`
  - Funzione: crea un array di grandezza `n` con ogni elemento inizializzato a `type`

- **int**
  - Parametri: `a`
  - Funzione: converte una stringa o un float `a` in un intero se possibile

- **float**
  - Parametri: `a`
  - Funzione: converte una stringa o un intero `a` in un float se possibile

- **string**
  - Parametri: `a`
  - Funzione: converte un elemento `a` in stringa

- **print**
  - Parametri: `a`
  - Funzione: Stampa `a` su console

- **println**
  - Parametri: `a`
  - Funzione: Stampa `a` su console andando a capo

- **read**
  - Parametri: Nessuno
  - Funzione: Legge da input

- **import**
  - Parametri: `path`
  - Funzione: Importa un modulo o un file dall'`path` specificato


