---
title: "Elastic Search Quickstart"
date: 2019-01-23T21:08:24+08:00
draft: false
tags: [ElasticSearch]
---

# Prerequisites
## What is Elasticsearch

- high-performance RESTful search engine and document store

## Start Elasticsearch

On Windows:
`elastic‑search‑dir\bin\elasticsearch.bat`

# Using cURL for REST commands

## Create an index

`$ curl ‑XPUT "http://localhost:9200/music/"`

## Insert a document

```
$ curl XPUT "http://localhost:9200/music/songs/1" ‑d '
{ "name": "Deck the Halls", "year": 1885, "lyrics": "Fa la la la la" }'
```

## View a document

```
$ curl ‑XGET "http://localhost:9200/music/songs/1"
```
Output:
```
{"_index":"music","_type":"songs","_id":"1","_version":1,"found":true,"_source":
{ "name": "Deck the Halls", "year": 1885, "lyrics": "Fa la la la la" }}
```
## Update a document

```
$ curl ‑XPUT "http://localhost:9200/music/lyrics/1" ‑d '{ "name":
"Deck the Halls", "year": 1886, "lyrics": "Fa la la la la" }'
```

## Delete a document
```
curl ‑XDELETE "http://localhost:9200/music/lyrics/1"
```

## Insert a document from a file
+ `$ curl ‑XPUT "http://localhost:9200/music/lyrics/2" ‑d @caseyjones.json`

```
{
  "artist": "Wallace Saunders",
  "year": 1909,
  "styles": ["traditional"],
  "album": "Unknown",
  "name": "Ballad of Casey Jones",
  "lyrics": "Come all you rounders if you want to hear
The story of a brave engineer
Casey Jones was the rounder's name....
Come all you rounders if you want to hear
The story of a brave engineer
Casey Jones was the rounder's name
On the six‑eight wheeler, boys, he won his fame
The caller called Casey at half past four
He kissed his wife at the station door
He mounted to the cabin with the orders in his hand
And he took his farewell trip to that promis'd land

Chorus:
Casey Jones‑‑mounted to his cabin
Casey Jones‑‑with his orders in his hand
Casey Jones‑‑mounted to his cabin
And he took his... land"
}
```
+ `curl ‑XPUT "http://localhost:9200/music/lyrics/3" ‑d @walking.json`
  
```

{
  "artist": "Clarence Ashley",
  "year": 1920
  "name": "Walking Boss",
  "styles": ["folk","protest"],
  "album": "Traditional",
  "lyrics": "Walkin' boss
Walkin' boss
Walkin' boss
I don't belong to you

I belong
I belong
I belong
To that steel driving crew

Well you work one day
Work one day
Work one day
Then go lay around the shanty two"
}
```

## Search the REST API

`curl ‑XGET "http://localhost:9200/music/lyrics/_search?q=lyrics:'you'"`

Response

```
(newline){"took":107,"timed_out":false,"_shards":{"total":5,"successful":5,"failed":0},"hits":{"total":2,"max(newline)_score":0.15625,"hits":[{"_index":"music","_type":"songs","_id":"2","_(newline)score":0.15625,"_source":{"artist": "Wallace Saunders","year": 1909,"styles":(newline)["traditional"],"album": "Unknown","name": "Ballad of Casey Jones","lyrics": "Come all you rounders(newline)if you want to hear The story of a brave engineer Casey Jones was the rounder's name.... Come all(newline)you rounders if you want to hear The story of a brave engineer Casey Jones was the rounder's name(newline)On the six‑eight wheeler, boys, he won his fame The caller called Casey at half past four He kissed(newline)his wife at the station door He mounted to the cabin with the orders in his hand And he took his(newline)farewell trip to that promis'd land Chorus: Casey Jones‑‑mounted to his cabin Casey Jones‑‑with his(newline)orders in his hand Casey Jones‑‑mounted to his cabin And he took his... land"(newline)}},{"_index":"music","_type":"songs","_id":"3","_score":0.06780553,"_source":{"artist": "Clarence(newline)Ashley","year": 1920,"name": "Walking Boss","styles": ["folk","protest"],"album":(newline)"Traditional","lyrics": "Walkin' boss Walkin' boss Walkin' boss I don't belong to you I belong I(newline)belong I belong To that steel driving crew Well you work one day Work one day Work one day Then go(newline)lay around the shanty two"}}]}}

```

## Use other comparators
`curl ‑XGET "http://localhost:9200/music/lyrics/_search?q=year:<1900`

## Restrict fields
`curl ‑XGET "http://localhost:9200/music/lyrics/_search?q=year:>1900&fields=year"`


## Examine the search return objects
```
{
    "took": 6,
    "timed_out": false,
    "_shards": {
        "total": 5,
        "successful": 5,
        "failed": 0
    },
    "hits": {
        "total": 2,
        "max_score": 1.0,
        "hits": {
            "_index": "music",
            "_type": "lyrics",
            "_id": "1",
            "_score": 1.0,
            "fields": {
                "year": 1920            
        }
        }, {
            "_index": "music",
            "_type": "lyrics",
            "_id": "3",
            "_score": 1.0,
            "fields": {
                "year": 1909            }
        }    
    }
}
```

## Use the JSON query DSL
```
{
    "query" : {
        "match" : {
            "album" : "Traditional"
        }
    }
}
```
`$ curl ‑XGET "http://localhost:9200/music/lyrics/_search" ‑d @query.json`

# Java API
