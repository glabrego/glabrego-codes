require 'pry'

class Bottles
  def verse(bottles)
    case
    when bottles > 2
      plural_verse(bottles)
    when bottles == 2
      two_verse(bottles)
    when bottles == 1
      one_verse
    else
      zero_verse
    end
  end

  def verses(initial_bottles, final_bottles)
    initial_bottles.downto(final_bottles).map {|i| verse(i)}.join("\n")
  end

  def song
    verses(99, 0)
  end

  private

  def plural_verse(bottles)
    "#{bottles} bottles of beer on the wall, #{bottles} bottles of beer.
Take one down and pass it around, #{bottles - 1} bottles of beer on the wall.\n"
  end

  def two_verse(bottles)
    "#{bottles} bottles of beer on the wall, #{bottles} bottles of beer.
Take one down and pass it around, #{bottles - 1} bottle of beer on the wall.\n"
  end

  def one_verse
    "1 bottle of beer on the wall, 1 bottle of beer.
Take it down and pass it around, no more bottles of beer on the wall.\n"
  end

  def zero_verse
    "No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.\n"
  end
end
