# ramses

Attingendo ai dati di IPDOM Ramses elabora un grafico delle ultime 10 ore per tutte le interfacce del device passato come argomento e identifica se negli ultimi 15 minuti ci sono stati dei valori che si sono discostati di oltre due deviazioni standard in più o in meno rispetto alla media mobile degli ultimi 20 punti.

Forse avrei dovuto mettere qualche virgola nel periodo di prima...

Enjoy! 

Dal vostro amichevole Gopher di quartiere.

Sintassi: 
    
    ramses -s=numero di deviazioni standard da considerare device da interrogare

Esempio:

    ramses -s=2.5 se-fi1-9


per funzionare serve settare delle variabili d'ambiente.

Su windows:

    set username=matricola
    set password=password di posta