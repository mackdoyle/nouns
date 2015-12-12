// ===========================================================================
// Nouns Graph
// /usr/local/Cellar/neo4j/2.3.0/libexec/data/graph.db
// ===========================================================================

// Structural Nodes
// ---------------------------------------------------------------------------

// Nouns
CREATE (person:noun { name : 'Person' }) RETURN "hello", person.name
CREATE (place:noun { name : 'Place' }) RETURN "hello", place.name
CREATE (thing:noun { name : 'Thing' }) RETURN "hello", thing.name
CREATE (concept:noun { name : 'Concept' }) RETURN "hello", concept.name


// Person Categories
CREATE (actor:person_category { name:'Actor' })
CREATE (author:person_category { name:'Author' })
CREATE (musician:person_category { name:'Musician' })
CREATE (technologist:person_category { name:'Technologist' })


// Place Categories
CREATE (bar:place_category { name:'Bar' })
CREATE (bar:place_category { name:'Coffee' })
CREATE (bar:place_category { name:'Dining' })
CREATE (bar:place_category { name:'Grocery' })
CREATE (bar:place_category { name:'Park' })
CREATE (trail:place_category { name:'Trail' })
CREATE (shopping:place_category { name:'Shopping' })

// Place Traits
CREATE (atmosphere:place_traits { name:'Atmosphere' })
CREATE (biking:place_traits { name:'Biking' })
CREATE (climbing:place_traits { name:'Climbing' })
CREATE (delivery:place_traits { name:'Delivery' })
CREATE (mixology:place_traits { name:'Mixology' })
CREATE (skiing:place_traits { name:'Skiing' })
CREATE (outdoor:place_traits { name:'Outdoor' })
CREATE (pets:place_traits { name:'Pets' })
CREATE (wifi:place_traits { name:'Wi-Fi' })


// Things Categories
CREATE (movie:thing_category { name:'Movie' })
CREATE (television_series:thing_category { name:'Television Series' })
CREATE (photography:thing_category { name:'Photography' })
CREATE (sporting_equipment:thing_category { name:'Sporting Equipment' })
CREATE (computers:thing_category { name:'Computers' })

// Concepts
// @TODO


// ===========================================================================
// LOCATIONS
// ===========================================================================

// Localities
// http://www.unece.org/fileadmin/DAM/cefact/locode/usp.htm
// ---------------------------------------------------------------------------
CREATE (beavertonor:locality { name:'Beaverton', locode: 'BVW', coordinates: ''  })
CREATE (bendor:locality { name:'Bend', locode: 'BZO', coordinates: '-121.18, 44.03'  })
CREATE (hoodriveror:locality { name:'Hood River', locode: 'HDX', coordinates: ''  })
CREATE (portlandor:locality { name:'Portland', locode: 'PDX', coordinates: ''  })
CREATE (mcminnvilleor:locality { name:'McMinnville', locode: 'MMC', coordinates: ''  })
CREATE (mthoodvillager:locality { name:'Mt. Hood Village', locode: '', coordinates: '-121.90, 45.35'  })
CREATE (tigardor:locality { name:'Tigard', locode: 'TID', coordinates: ''  })


// Regions
// ---------------------------------------------------------------------------
CREATE (oregon:region { name:'Oregon', recode: 'OR', coordinates: ''  })

// Countries
// ---------------------------------------------------------------------------
CREATE (oregon:country { name:'United States', cocode: 'US', coordinates: ''  })

// ===========================================================================
// NOUNS
// ===========================================================================

// Places
CREATE (powellscityofbooks1:place {
  name:'Powell\'s City of Books',
  street_address:'1005 W Burnside St',
  extended_address:'',
  locality:'Portland',
  region:'Oregon',
  postal_code:'',
  country_code:'US',
  phone_number:'5032284651',
  coordinates:'-122.6814431,45.5231094',

  link:'powells.com',
  image:''
})
CREATE (powellscityofbooks1)-[:HAS_TRAIT]->(wifi)
CREATE (powellscityofbooks1)-[:IS_TYPE]->(Shopping)



CREATE (philip:Person {name:"Philip"})-[:IS_FRIEND_OF]->(emil:Person {name:"Emil"}),
       (philip)-[:IS_FRIEND_OF]->(michael:Person {name:"Michael"}),
       (philip)-[:IS_FRIEND_OF]->(andreas:Person {name:"Andreas"})
create (sushi:Cuisine {name:"Sushi"}), (nyc:City {name:"New York"}),
       (iSushi:Restaurant {name:"iSushi"})-[:SERVES]->(sushi),(iSushi)-[:LOCATED_IN]->(nyc),
       (michael)-[:LIKES]->(iSushi),
       (andreas)-[:LIKES]->(iSushi),
       (zam:Restaurant {name:"Zushi Zam"})-[:SERVES]->(sushi),(zam)-[:LOCATED_IN]->(nyc),
       (andreas)-[:LIKES]->(zam)
