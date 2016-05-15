require 'json-schema'
class ItemSchema
  include Mongoid::Document
  validates :collection_name, presence: true, length: { minimum: 3 }
  validate :fully_validate_item_schema

  field :properties, type: Hash
  field :schema, type: String
  field :type, type: String
  field :additional_properties, type: Mongoid::Boolean
  field :description, type: String
  field :collection_name, type: String
  field :required, type: Array

  def to_json_schema
    schema = attributes.stringify_keys
    schema['$schema'] = schema.delete 'schema'
    schema['additionalProperties'] = schema.delete 'additional_properties'
    schema['collectionName'] = schema.delete 'collection_name'
    schema['id'] = schema.delete('_id').to_s
    schema
  end

  private
  def fully_validate_item_schema
    schema = attributes.stringify_keys
    if JSON::Validator.fully_validate_schema(schema).empty?

    end
  end

end
