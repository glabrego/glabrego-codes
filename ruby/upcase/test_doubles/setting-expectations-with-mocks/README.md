# Test Doubles / Setting Expectations With Mocks

Hey there! We're [thoughtbot](https://thoughtbot.com), a design and
development consultancy that brings your digital product ideas to life.
We also love to share what we learn.

This coding exercise comes from [Upcase](https://thoughtbot.com/upcase),
the online learning platform we run. It's part of the
[Test Doubles](https://thoughtbot.com/upcase/test-doubles) course and is just one small sample of all
the great material available on Upcase, so be sure to visit and check out the rest.

## Exercise Intro

You've already learned how to create and use stubs with the `double` and `allow` methods. Stubs can be used to unit test objects which use query methods from other objects:

``` ruby
describe Signup do
  context "#email" do
    it "delegates to the user" do
      user = double("user")
      allow(user).to receive(:email).and_return("user@example.com")
      signup = Signup.new(user: user)

      expect(signup.email).to eq(user.email)
    end
  end
end
```

In this example, we're testing that the `Signup` class uses the value from the `email` method on `User`. We call `email` a "query method," because it asks a question ("hey user, what is your email?") but doesn't tell the user to do anything.

In real life, however, your objects will also have command methods: methods which mutate or perform other side effects. When testing these methods, stubs will not suffice.

Here's an example of attempting to use a stub to test that we're invoking the `update_attributes` method on a user:

``` ruby
describe Signup do
  context "#update_attributes" do
    it "sets attributes on the user" do
      user = double("user")
      allow(user).to receive(:update_attributes)
      signup = Signup.new(user: user)

      signup.update_attributes(email: "new_user@example.com")
    end
  end
end
```

Although the test will pass if `Signup` correctly calls `update_attributes`, it will also pass if `update_attributes` is never called at all. This is where a mock comes in.

A mock is just like a stub, except that it doesn't just allow methods to be invoked; it expects it.

Here's the same test using a mock:

``` ruby
describe Signup do
  context "#update_attributes" do
    it "sets attributes on the user" do
      user = double("user")
      expect(user).to receive(:update_attributes).
        with(email: "new_user@example.com")
      signup = Signup.new(user: user)

      signup.update_attributes(email: "new_user@example.com")
    end
  end
end
```

Mocks are auto-verifying. You set the expectation when mocking out the method, and RSpec will automatically verify that all mocked methods were invoked as expected at the end of the test. If the `update_attributes` method is never invoked, the test will fail.

In this exercise, you'll use mocks to verify that a command method is invoked as expected.

## Instructions

To start, you'll want to clone and run the setup script for the repo

    git clone git@github.com:thoughtbot-upcase-exercises/setting-expectations-with-mocks.git
    cd setting-expectations-with-mocks
    bin/setup

After running `bin/setup`, edit `spec/signup_spec.rb` and refactor the tests to use mocks instead of hitting the database.

Use `expect` when necessary to test command methods, and use `allow` when you only need to use the return value.

## Tips and Tricks

Useful links:

- Check out our Weekly Iteration episode on [Stubs, Mocks, Spies, and Fakes](https://upcase.com/videos/stubs-mocks-spies-and-fakes)
- Check out the rspec-mocks summary of [message expectations](https://github.com/rspec/rspec-mocks#message-expectations)

## Featured Solution

Check out the [featured solution branch](https://github.com/thoughtbot-upcase-exercises/setting-expectations-with-mocks/compare/featured-solution#toc) to
see the approach we recommend for this exercise.

## Forum Discussion

If you find yourself stuck, be sure to check out the associated
[Upcase Forum discussion](https://forum.upcase.com/t/test-doubles-setting-expectations-with-mocks/4611)
for this exercise to see what other folks have said.

## Next Steps

When you've finished the exercise, head on back to the
[Test Doubles](https://thoughtbot.com/upcase/test-doubles) course to find the next exercise,
or explore any of the other great content on
[Upcase](https://thoughtbot.com/upcase).

## License

setting-expectations-with-mocks is Copyright © 2015-2018 thoughtbot, inc. It is free software,
and may be redistributed under the terms specified in the
[LICENSE](/LICENSE.md) file.

## Credits

![thoughtbot](https://presskit.thoughtbot.com/assets/images/logo.svg)

This exercise is maintained and funded by
[thoughtbot, inc](http://thoughtbot.com/community).

The names and logos for Upcase and thoughtbot are registered trademarks of
thoughtbot, inc.
