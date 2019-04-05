# ramses

[![Build Status](https://travis-ci.org/axamon/ramses.svg?branch=master)](https://travis-ci.org/axamon/ramses)
[![Maintainability](https://api.codeclimate.com/v1/badges/55cbc6bd7cdf6afd7c52/maintainability)](https://codeclimate.com/github/axamon/ramses/maintainability)


Attingendo ai dati di IPDOM Ramses elabora un grafico delle ultime 10 ore per tutte le interfacce del device passato come argomento e identifica se negli ultimi 15 minuti ci sono stati dei valori che si sono discostati di oltre due deviazioni standard in più o in meno rispetto alla media mobile degli ultimi 20 punti.

Ramses inoltre monitora un sottoinsieme dei NAS esistenti ogni 5 minuti e allerta via mail se ci 
sono scostamenti del numero di sessioni ppp molto elvati verso il basso.

Per configure l'applicativo bisogna compilare i campi di un file json di configurazione:

``` json
{
    "IPDOMUser": "user",
    "IPDOMPassword": "fdsfdsf",
    "IPDOMSnmpReceiver": "127.0.0.1",
    "IPDOMSnmpPort": 162,
    "IPDOMSnmpCommunity": "public",
    "NasDaIgnorare": "nasDaIgnorare.json",
    "NasInventory": "nasInventory.json",
    "Sigma": 2.5,
    "Soglia": 0.1,
    "URLSessioniPPP": "https://ipw.telecomitalia.it/ipwmetrics/api/v1/rawmetrics/kpi.ppoe.slot?device=",
    "URLTail7d": "&start=7d-ago&end=5m-ago&aggregator=sum",
    "SmtpPort": 587,
    "SmtpServer": "smtp.gmail.com",
    "SmtpUser": "user@sender.com",
    "SmtpPassword": "fsdfsdfdsfds",
    "SmtpSender": "user@domain.com",
    "SmtpFrom": "user@domain.com",
    "SmtpTo": "receiver@domain2.com"
}
```

# Sintassi

    ramses filediconfigurazione.json
