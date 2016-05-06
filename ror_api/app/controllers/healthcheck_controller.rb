class HealthcheckController < ApplicationController

  def index
    render json: { status: "working" }
  end

end
