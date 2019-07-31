# -*- coding: utf-8 -*-
import re

import scrapy

from freelance.items import TaskItem, PRICE_PERIOD_HOURLY, PRICE_PERIOD_PROJECT


re_price_amount = re.compile(r'\D*')


class FreelansimSpider(scrapy.Spider):
    name = 'freelansim'
    allowed_domains = ['freelansim.ru']
    start_urls = ['https://freelansim.ru/']

    def parse(self, response):
        for dom_item in response.css('.content-list__item'):
            task_url = response.urljoin(dom_item.css('.task__title a::attr(href)').extract_first())
            task_title = dom_item.css('.task__title a::text').extract_first()
            task_views = self.get_views(dom_item)
            task_responses = self.get_responses(dom_item)
            task_is_safe_deal = len(dom_item.css('.icon_safe_deal')) > 0
            task_tags = dom_item.css('.tags__item_link::text').extract()
            task_price_negotiated = len(dom_item.css('.negotiated_price')) > 0
            task_price_amount = self.get_price_amount(dom_item)
            task_price_period = self.get_price_period(dom_item)

            item = TaskItem(
                url=task_url,
                title=task_title,
                views=task_views,
                responses=task_responses,
                is_safe_deal=task_is_safe_deal,
                tags=task_tags,
                price_negotiated=task_price_negotiated,
                price_amount=task_price_amount,
                price_period=task_price_period
            )

            yield scrapy.Request(url=task_url, callback=self.parse_detail, meta={"item": item})

    def parse_detail(self, response):
        item = response.meta.get('item')
        item['description'] = response.css('.task__description').xpath('text()').extract_first().strip()

        yield item

    @staticmethod
    def get_views(dom_item):
        try:
            return int(dom_item.css('.params__views > .params__count::text').extract_first())
        except (ValueError, TypeError):
            return 0

    @staticmethod
    def get_responses(dom_item):
        try:
            return int(dom_item.css('.params__responses > .params__count::text').extract_first())
        except (ValueError, TypeError):
            return 0

    @staticmethod
    def get_price_amount(dom_item):
        try:
            return int(re_price_amount.sub('', dom_item.css('.task__price > .count::text').extract_first()))
        except (ValueError,TypeError):
            return 0

    @staticmethod
    def get_price_period(dom_item):
        suffix = dom_item.css('.task__price > .count > .suffix::text').extract_first()
        
        if suffix == 'за час':
            return PRICE_PERIOD_HOURLY
        elif suffix == 'за проект':
            return PRICE_PERIOD_PROJECT
        else:
            return None
