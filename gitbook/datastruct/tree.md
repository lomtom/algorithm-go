# 二叉树

# 知识点

## 二叉树遍历

**前序遍历**：**先访问根节点**，再前序遍历左子树，再前序遍历右子树 **中序遍历**：先中序遍历左子树，**再访问根节点**，再中序遍历右子树 **后序遍历**：先后序遍历左子树，再后序遍历右子树，**再访问根节点**

注意点

- 以根访问顺序决定是什么遍历
- 左子树都是优先右子树

<article class="message message-immersive is-primary">
<div class="message-body">
<i class="fas fa-globe-americas mr-2"></i>本文在：<a href="https://blog.csdn.net/qq_41929184/article/details/116799994">csdn</a>、<a href="https://mp.weixin.qq.com/s/HmMw8sVKEMcsdXL1im0-Lw">公众号</a>同步更新。
</div>
</article>



<escape><!-- more --></escape>



### 你能get到的知识点

1. 抽象工厂模式的介绍
2. 抽象工厂模式通过代码的实现

> 本文首发于CSDN，作者：lomtom
>
> 原文链接：[https://blog.csdn.net/qq_41929184/article/details/118366444](https://blog.csdn.net/qq_41929184/article/details/118366444)
>
> 个人网站：[https://lomtom.top](https://lomtom.top)，个人公众号：[博思奥园](https://mp.weixin.qq.com/s/HmMw8sVKEMcsdXL1im0-Lw)，同步更新。
>
> 你的支持就是我最大的动力。



设计模式系列：

1. [设计模式（一）工厂模式](https://blog.csdn.net/qq_41929184/article/details/117955365)
2. [设计模式（二）抽象工厂模式](https://blog.csdn.net/qq_41929184/article/details/118338143)





## 简介

这一节主要讲解的是创建者模式中的抽象工厂模式

抽象工厂模式是一种创建型设计模式， 它能创建一系列相关的对象， 而无需指定其具体类。![图片来自https://refactoringguru.cn](https://img-blog.csdnimg.cn/20210629181538373.png)

- 适用场景

1. 如果代码需要与多个不同系列的相关产品交互， 但是由于无法提前获取相关信息， 或者出于对未来扩展性的考虑， 你不希望代码基于产品的具体类进行构建， 在这种情况下， 你可以使用抽象工厂。
2. 如果你有一个基于一组抽象方法的类， 且其主要功能因此变得不明确， 那么在这种情况下可以考虑使用抽象工厂模式。

## 基于抽象工厂模式的应用

简单步骤：

1. 为所有产品声明抽象产品接口。 然后让所有具体产品类实现这些接口。
2. 声明抽象工厂接口， 并且在接口中为所有抽象产品提供一组构建方法。
3. 为每种产品变体实现一个具体工厂类。
4. 在应用程序中开发初始化代码。 该代码根据应用程序配置或当前环境， 对特定具体工厂类进行初始化。 然后将该工厂对象传递给所有需要创建产品的类。
5. 找出代码中所有对产品构造函数的直接调用， 将其替换为对工厂对象中相应构建方法的调用。

> 场景升级实例： 	小葛现在不仅是参加Jd平台的抽奖活动，也同样参加淘宝的抽奖活动，奖品和之前一样，同样有三种，分别为300元购物券、Iphone12 和3000元现金


工程结构

```java
└─com
    └─lomtom
        └─demo_0_5
            │  Test.java
            │
            ├─factory
            │  │  AwardFactory.java
            │  │  FactoryFactory.java
            │  │
            │  └─impl
            │      ├─jd
            │      │      JdAwardFactory.java
            │      │
            │      └─taobao
            │              TaobaoAwardFactory.java
            │
            └─service
                │  AwardService.java
                │
                └─impl
                    ├─jd
                    │      CashAwardService.java
                    │      IphoneAwardService.java
                    │      MallCardAwardService.java
                    │
                    └─taobao
                            CashAwardService.java
                            IphoneAwardService.java
                            MallCardAwardService.java
```

![](https://img-blog.csdnimg.cn/20210630180555962.png#pic_center)

1. 为所有产品声明抽象产品接口。 然后让所有具体产品类实现这些接口。

```java
1.发放奖品的接口
public interface AwardService {
    void getAward(String username);
}

2.编写具体的实现类实现真正的奖品发放
jd发放奖品实现类
public class CashAwardService implements AwardService {
    private String award = "jd     ---- 3000元 现金";
    @Override
    public void getAward(String username){
        System.out.println(username + "获得了" + award);
    }
}

public class IphoneAwardService implements AwardService {
    private String award = "jd     ---- Iphone 12";
    @Override
    public void getAward(String username){
        System.out.println(username + "获得了" + award);
    }
}

public class MallCardAwardService implements AwardService {
    private String award = "jd     ---- 1000元 购物卡";
    @Override
    public void getAward(String username){
        System.out.println(username + "获得了" + award);
    }
}


taobao发放奖品实现类
public class CashAwardService implements AwardService {
    private String award = "taobao ---- 3000元 现金";
    @Override
    public void getAward(String username){
        System.out.println(username + "获得了" + award);
    }
}

public class IphoneAwardService implements AwardService {
    private String award = "taobao ---- Iphone 12";
    @Override
    public void getAward(String username){
        System.out.println(username + "获得了" + award);
    }
}

public class MallCardAwardService implements AwardService {
    private String award = "taobao ---- 1000元 购物卡";
    @Override
    public void getAward(String username){
        System.out.println(username + "获得了" + award);
    }
}
```

2. 声明抽象工厂接口， 并且在接口中为所有抽象产品提供一组构建方法。

```java
声明抽象创建者，通过一个标志来决定创建哪一个工厂
public class FactoryFactory {
    public static AwardFactory createAwardFactory(Integer factoryNumber) {
        if (factoryNumber == 1){
            return new JdAwardFactory();
        }else{
            return new TaobaoAwardFactory();
        }
    }
}
```

3. 为每种产品变体实现一个具体工厂类。

```java
1.声明具体创建者
public interface AwardFactory {
    AwardService getAwardService(Integer awardNumber);
}

2.实现每一个工厂实现具体的创建
jd工厂，注意这里引入的是jd的发放接口（因为名字设为一样，只是以包将jd和淘宝区别开）
import com.lomtom.demo_0_5.factory.AwardFactory;
import com.lomtom.demo_0_5.service.AwardService;
import com.lomtom.demo_0_5.service.impl.jd.CashAwardService;
import com.lomtom.demo_0_5.service.impl.jd.IphoneAwardService;
import com.lomtom.demo_0_5.service.impl.jd.MallCardAwardService;

/**
 * @author lomtom
 **/
public class JdAwardFactory implements AwardFactory {

    @Override
    public AwardService getAwardService(Integer awardNumber) {
        if (awardNumber == 1){
            return new MallCardAwardService();
        }else if (awardNumber == 2){
            return new IphoneAwardService();
        }else{
            return new CashAwardService();
        }
    }
}

taobao工厂
import com.lomtom.demo_0_5.service.AwardService;
import com.lomtom.demo_0_5.service.impl.taobao.CashAwardService;
import com.lomtom.demo_0_5.service.impl.taobao.IphoneAwardService;
import com.lomtom.demo_0_5.service.impl.taobao.MallCardAwardService;

/**
 * @author lomtom
 **/
public class TaobaoAwardFactory implements com.lomtom.demo_0_5.factory.AwardFactory {

    @Override
    public AwardService getAwardService(Integer awardNumber) {
        if (awardNumber == 1){
            return new MallCardAwardService();
        }else if (awardNumber == 2){
            return new IphoneAwardService();
        }else{
            return new CashAwardService();
        }
    }
}
```

4. 在应用程序中开发初始化代码。 该代码根据应用程序配置或当前环境， 对特定具体工厂类进行初始化。 然后将该工厂对象传递给所有需要创建产品的类。

```java
进行测试验证
public class Test {

    public static void main(String[] args) {
        int index = 0;
        while(index++ < 10) {
            String employee = "小葛";
            Random random = new Random();
            //小葛随机去平台进行抽奖
            Integer factoryNumber = random.nextInt(2);
            //小葛随机获得奖品
            Integer awardNumber = random.nextInt(3);
            System.out.print(employee + "抽奖兑换------    ");

            AwardFactory factory = FactoryFactory.createAwardFactory(factoryNumber);
            AwardService service = factory.getAwardService(awardNumber);
            service.getAward(employee);

        }

    }
}
```

结果：

```java
葛抽奖兑换------    小葛获得了jd     ---- 3000元 现金
小葛抽奖兑换------    小葛获得了taobao ---- Iphone 12
小葛抽奖兑换------    小葛获得了taobao ---- 3000元 现金
小葛抽奖兑换------    小葛获得了jd     ---- Iphone 12
小葛抽奖兑换------    小葛获得了jd     ---- Iphone 12
小葛抽奖兑换------    小葛获得了taobao ---- Iphone 12
小葛抽奖兑换------    小葛获得了jd     ---- 1000元 购物卡
小葛抽奖兑换------    小葛获得了jd     ---- Iphone 12
小葛抽奖兑换------    小葛获得了taobao ---- 3000元 现金
小葛抽奖兑换------    小葛获得了jd     ---- 3000元 现金
```

使用抽象工厂，有比较多的**优点**：

* 可以确保同一工厂生成的产品相互匹配、避免客户端和具体产品代码的耦合。
* 符合`单一职责原则`和`开闭原则`。 可以将产品生成代码抽取到同一位置，使得代码易于维护； 向抽奖活动中引入新平台时，无需修改客户端代码。

当然，他也不是没有**缺点**的：当我们使用抽象工厂时，会引入其他的类与接口，这样会造成代码的复杂变高，理解更为复杂。
