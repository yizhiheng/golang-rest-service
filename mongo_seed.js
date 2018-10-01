db = db.getSiblingDB('golang-rest-service-db')

db.resources.insert({
  '_id': new ObjectId(),
  'title': 'Best title',
  'number': 100
})
