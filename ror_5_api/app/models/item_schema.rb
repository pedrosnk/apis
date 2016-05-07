class ItemSchema
  include Mongoid::Document
  field :properties, type: Hash
  field :schema, type: String
  field :type, type: String
  field :aditionalProperties, type: Mongoid::Boolean
  field :description, type: String
  field :collection_name, type: String
end
