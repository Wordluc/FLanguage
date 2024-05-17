# FLanguage

Benvenuto su FLanguage! Questo progetto è stato creato come esperimento per comprendere il funzionamento degli interpreti e per imparare il linguaggio di programmazione Go.
Puoi trovare un esempio di codice Flanguage in "helloWord.txt".
- **RunFile**: Per eseguire il codice Flanguage bisogna spostarsi in ./bin/ e eseguire  `.\FLanguage.exe "<file.txt>"`
- **RunRepl**: Per eseguire il repl di Flanguage bisogna spostarsi in ./bin/ e eseguire  `.\FLanguage.exe "r"`

## Caratteristiche Principali

- **Funzioni**: Puoi definire e utilizzare funzioni per organizzare il tuo codice in moduli riutilizzabili e per gestire compiti specifici.Le funzioni non possono modificare l'ambiente esterno in cui sono definite.
  
- **Array**: Manipola collezioni di dati con facilità utilizzando array, consentendo di memorizzare e accedere a elementi in modo efficiente.
  
- **Hashtable**: Utilizza hashtable per implementare strutture dati chiave-valore, ideali per la gestione di associazioni tra dati e inline function.
  
- **Istruzioni Condizionali**: Controlla il flusso del programma utilizzando istruzioni condizionali if-else, consentendo di eseguire operazioni diverse in base a determinate condizioni.
  
- **Cicli While**: Itera attraverso i dati o esegui operazioni ripetute finché una condizione specifica è vera utilizzando i cicli while.
  
- **Funzioni Inline**: Definisci funzioni direttamente nel contesto del codice principale per migliorare la leggibilità e la modularità del codice.
  
- **Importazione di Codice Esterno**: Importa facilmente codice da altri file per organizzare il tuo progetto in moduli separati e riutilizzare il codice esistente,i moduli devono essere composti esclusivamente da funzioni. Per richiamare una funzione, è necessario seguire una regola specifica: si concatena il nome del modulo (senza estensione) con il nome della funzione, separati da un underscore (_).
  
- **Oggetti**:Vi è la possibilità di poter creare un hashtable con al interno delle innerfunction che potranno interaggire con la hashtable attraverso la parola chiave "this"
- **Fine File**:Bisogna indicare la fine del file mediante le keyword  `END`

## Sintassi

- **let**: Utilizzata per dichiarare variabili.Una volta assegnato un elemento non è possibile cambiare il tipo della variabile
  - Esempio: `let a = 2;`

- **import**: Utilizzato per ottenere un modulo.
  - Esempio:
    ```
    let tree=import("BinarySearch.txt");`
    tree{"Run"}([1,2,3,4,7],4);
    ```
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

- **import()**: Utilizzata per ritornare il modulo sotto forma di oggetto.
  - Esempio: `let modulo=import("nome_modulo");`

- **this{}**: Utilizzata per fare riferimento all'hashtable corrente all'interno di un contesto di programmazione orientato agli oggetti.
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
  - Funzione: restituisce il numero di elementi di un array `a` o il numero di caratteri di una stringa `a`

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
  - Funzione: Ottiene un modulo dall'`path` e lo restituisce come oggetto


