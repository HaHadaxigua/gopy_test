# This is a sample Python script.

# Press Shift+F10 to execute it or replace it with your code.
# Press Double Shift to search everywhere for classes, files, tool windows, actions, and settings.


from factory_out import factory

# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    objA = factory.BuildXInterface("A")
    objA.DoA()

    objB = factory.BuildXInterface("B")
    objB.DoB()

    objs = factory.BuildXInterfaces(["A", "B"]) # not work


# See PyCharm help at https://www.jetbrains.com/help/pycharm/
