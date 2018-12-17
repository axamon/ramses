# ramses

Attingendo ai dati di IPDOM Ramses elabora un grafico delle ultime 10 ore per tutte le interfacce del device passato come argomento e identifica se negli ultimi 15 minuti ci sono stati dei valori che si sono discostati di oltre due deviazioni standard in più o in meno rispetto alla media mobile degli ultimi 20 punti.

Ramses inoltre monitora un sottoinsieme dei NAS esistenti ogni 5 minuti e allerta via mail se ci 
sono scostamenti del numero di sessioni ppp molto elvati verso il basso.


per configure l'applicativo bisogna compilare i campi di un file json di configurazione:

{
    "Sigma": 2.5,
    "IPDOMUser": "00246506",
    "IPDOMPassword": "fdfsfsfsdfsd",
    "NasInventory": "nasInventory.json",
    "URLSessioniPPP": "https://ipw.telecomitalia.it/ipwmetrics/api/v1/rawmetrics/kpi.ppoe.slot?device=",
    "URLTail7d": "&start=7d-ago&end=5m-ago&aggregator=sum",
    "SmtpPort": 587,
    "SmtpServer": "smtp.gmail.com",
    "SmtpUser": "alberto.bregliano@gmail.com",
    "SmtpPassword": "fdsfdsfsdf",
    "SmtpSender": "alberto.bregliano@gmail.com",
    "SmtpFrom": "alberto.bregliano@gmail.com",
    "SmtpTo": "alberto.bregliano@protonmail.com"

}

Sintassi:

    ramses filediconfigurazione.json


Enjoy!  

Dal vostro amichevole Gopher di quartiere.

