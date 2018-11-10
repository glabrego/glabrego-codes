require "account"
require "user"

class Signup
  attr_reader :user

  def initialize(account_name:, email:)
    @account_name = account_name
    @email = email
  end

  def save
    account = Account.create!(name: @account_name)
    @user = User.create!(account: account, email: @email)
    true
  end
end
