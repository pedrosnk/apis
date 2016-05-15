class ItemSchemasController < ApplicationController
  before_action :set_item_schema, only: [:show, :update, :destroy]

  # GET /item_schemas
  def index
    @item_schemas = ItemSchema.all

    render json: @item_schemas
  end

  # GET /item_schemas/1
  def show
    render json: @item_schema
  end

  # POST /item_schemas
  def create
    @item_schema = ItemSchema.new(item_schema_params)

    if @item_schema.save
      render json: @item_schema, status: :created, location: @item_schema
    else
      render json: @item_schema.errors, status: :unprocessable_entity
    end
  end

  # PATCH/PUT /item_schemas/1
  def update
    if @item_schema.update(item_schema_params)
      render json: @item_schema
    else
      render json: @item_schema.errors, status: :unprocessable_entity
    end
  end

  # DELETE /item_schemas/1
  def destroy
    @item_schema.destroy
  end

  private
    # Use callbacks to share common setup or constraints between actions.
    def set_item_schema
      @item_schema = ItemSchema.find(params[:id])
    end

    # Only allow a trusted parameter "white list" through.
    def item_schema_params
      params.permit(:properties, :schema, :type, :additionalProperties,
                    :description, :collection_name, :required)

    end
end
