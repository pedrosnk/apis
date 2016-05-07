class ItemSchema
  include Mongoid::Document
  validate :fully_validate_item_schema

  field :properties, type: Hash
  field :schema, type: String
  field :type, type: String
  field :aditionalProperties, type: Mongoid::Boolean
  field :description, type: String
  field :collection_name, type: String
  field :required, type: Array

  private
  def fully_validate_item_schema
    JSON::Validator.fully_validate_schema(attributes).empty?
  end
end
