# Careerjet Go
Unofficial Careerjet API Golang Client Library

## Apply for Affiliate ID.
- Requires to open a Careerjet partner account: http://www.careerjet.co.uk/partners

## Mandatory Search Params

* `affid`      : Affiliate ID provided by Careerjet. Requires to open a Careerjet partner account (http://www.careerjet.co.uk/partners).
* `user_ip`    : IP address of the end-user to whom the search results will be displayed.  
* `user_agent` : User agent of the end-user's browser.
* `url`        : URL of page that will display the search results

## Search Params

Please note that each parameter is optional.

* `keywords`: Keywords to match the title, content or company name of a job offer

* `location`: Location of requested jobs

* `sort`: Sort type. This can be: `relevance` (default) — sorted by decreasing relevancy, `date` — sorted by decreasing date and `salary` — sorted by decreasing salary.

* `start_num`: Position of returned job postings within the entire result space. Should be >= 1 and <= Number of hits.

* `pagesize`: Number of jobs returned in one call.

* `page`: Page number of returned jobs within the entire result space. Should be >=1. If this value is set, it overrides `start_num`.

* `contracttype`: Selected contract type.`p` — permanent job, `c` — contract, `t` — temporary, `i` — training, `v` — voluntary, none — all contract types.

* `contractperiod`: Selected contract period. `f` — full time, `p` — part time, none — all contract periods.

## Locale code

The locale code needs to be supplied in the contructor of the API client. It defines the default location as well as the language in
which the search results are returned. Each locale corresponds to a Careerjet site.

The default is 'en_GB'.

Available locale codes are:

```ruby 
    :cs_CZ  => 'http://www.careerjet.cz'                  ,
    :da_DK  => 'http://www.careerjet.dk'                  ,
    :de_AT  => 'http://www.careerjet.at'                  ,
    :de_CH  => 'http://www.careerjet.ch'                  ,
    :de_DE  => 'http://www.careerjet.de'                  ,
    :en_AE  => 'http://www.careerjet.ae'                  ,
    :en_AU  => 'http://www.careerjet.com.au'              ,
    :en_BD  => 'http://www.careerjet.com.bd'              ,
    :en_CA  => 'http://www.careerjet.ca'                  ,
    :en_CN  => 'http://www.careerjet.com.cn'              ,
    :en_HK  => 'http://www.careerjet.hk'                  ,
    :en_IE  => 'http://www.careerjet.ie'                  ,
    :en_IN  => 'http://www.careerjet.co.in'               ,
    :en_KW  => 'http://www.careerjet.com.kw'              ,
    :en_MY  => 'http://www.careerjet.com.my'              ,
    :en_NZ  => 'http://www.careerjet.co.nz'               ,
    :en_OM  => 'http://www.careerjet.com.om'              ,
    :en_PH  => 'http://www.careerjet.ph'                  ,
    :en_PK  => 'http://www.careerjet.com.pk'              ,
    :en_QA  => 'http://www.careerjet.com.qa'              ,
    :en_SG  => 'http://www.careerjet.sg'                  ,
    :en_GB  => 'http://www.careerjet.co.uk'               ,
    :en_US  => 'http://www.careerjet.com'                 ,
    :en_ZA  => 'http://www.careerjet.co.za'               ,
    :en_SA  => 'http://www.careerjet-saudi-arabia.com'    ,
    :en_TW  => 'http://www.careerjet.com.tw'              ,
    :en_VN  => 'http://www.careerjet.vn'                  ,
    :es_AR  => 'http://www.opcionempleo.com.ar'           ,
    :es_BO  => 'http://www.opcionempleo.com.bo'           ,
    :es_CL  => 'http://www.opcionempleo.cl'               ,
    :es_CO  => 'http://www.opcionempleo.com.co'           ,
    :es_CR  => 'http://www.opcionempleo.co.cr'            ,
    :es_DO  => 'http://www.opcionempleo.com.do'           ,
    :es_EC  => 'http://www.opcionempleo.ec'               ,
    :es_ES  => 'http://www.opcionempleo.com'              ,
    :es_GT  => 'http://www.opcionempleo.com.gt'           ,
    :es_MX  => 'http://www.opcionempleo.com.mx'           ,
    :es_PA  => 'http://www.opcionempleo.com.pa'           ,
    :es_PE  => 'http://www.opcionempleo.com.pe'           ,
    :es_PR  => 'http://www.opcionempleo.com.pr'           ,
    :es_PY  => 'http://www.opcionempleo.com.py'           ,
    :es_UY  => 'http://www.opcionempleo.com.uy'           ,
    :es_VE  => 'http://www.opcionempleo.com.ve'           ,
    :fi_FI  => 'http://www.careerjet.fi'                  ,
    :fr_CA  => 'http://www.option-carriere.ca'            ,
    :fr_BE  => 'http://www.optioncarriere.be'             ,
    :fr_CH  => 'http://www.optioncarriere.ch'             ,
    :fr_FR  => 'http://www.optioncarriere.com'            ,
    :fr_LU  => 'http://www.optioncarriere.lu'             ,
    :fr_MA  => 'http://www.optioncarriere.ma'             ,
    :hu_HU  => 'http://www.careerjet.hu'                  ,
    :it_IT  => 'http://www.careerjet.it'                  ,
    :ja_JP  => 'http://www.careerjet.jp'                  ,
    :ko_KR  => 'http://www.careerjet.co.kr'               ,
    :nl_BE  => 'http://www.careerjet.be'                  ,
    :nl_NL  => 'http://www.careerjet.nl'                  ,
    :no_NO  => 'http://www.careerjet.no'                  ,
    :pl_PL  => 'http://www.careerjet.pl'                  ,
    :pt_PT  => 'http://www.careerjet.pt'                  ,
    :pt_BR  => 'http://www.careerjet.com.br'              ,
    :ru_RU  => 'http://www.careerjet.ru'                  ,
    :ru_UA  => 'http://www.careerjet.com.ua'              ,
    :sv_SE  => 'http://www.careerjet.se'                  ,
    :sk_SK  => 'http://www.careerjet.sk'                  ,
    :tr_TR  => 'http://www.careerjet.com.tr'              ,
    :uk_UA  => 'http://www.careerjet.ua'                  ,
    :vi_VN  => 'http://www.careerjet.com.vn'              ,
    :zh_CN  => 'http://www.careerjet.cn'                  ,
```

## Sample cURL request
```sh
~$  curl -X GET -H "user-agent: careerjet-api-client-v3.0.1" -H "referer: http://public.api.careerjet.net" "http://public.api.careerjet.net/search?affid=<your affid>&user_ip=<your user ip address>&user_agent=<your user agent>&locale_code=<locale code>&location=<country>&keywords=<keywords>"
```

### Take note
- **affid**: Affiliate ID
- **ip address**: to be encoded with `%2E` instead of `.`
- **keywords**: if multiple, to be encoded with `%2C` instead of `,`

## Reference
- Official Careerjet Python Client: https://github.com/careerjet/careerjet-api-client-python