# -*- coding: utf-8 -*-
import re

import scrapy

from freelance.items import TaskItem, PRICE_PERIOD_HOURLY, PRICE_PERIOD_PROJECT


re_task_id = re.compile(r'\/(\d+)\/')


class FlSpider(scrapy.Spider):
    name = 'fl'
    allowed_domains = ['fl.ru']
    start_urls = ['https://www.fl.ru/projects/']

    def parse(self, response):
        for dom_item in response.css('.b-post'):
            task_url = response.urljoin(dom_item.css('.b-post__link::attr(href)').extract_first())
            task_title = dom_item.css('.b-post__link::text').extract_first()

            item = TaskItem(
                url=task_url,
                title=task_title,
            )

            yield scrapy.Request(url=task_url, callback=self.parse_detail, meta={"item": item})

    def parse_detail(self, response):
        task_id = self.get_id(response)
        item = response.meta.get('item')

        item['description'] = self.get_description(response, task_id)

        yield item


    def get_description(self, response, id):
        try:
            return response.css('#projectp' + str(id)).extract_first().strip()
        except (AttributeError) as e:
            self.logger.warn('get description error (%s) %s' % (response.url, e))
            return ''

    @staticmethod
    def get_id(response):
        return re.search(r'/(\d+)/', response.url).group(1)